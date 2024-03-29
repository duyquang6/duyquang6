// Create README profile
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	firebase "firebase.google.com/go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	MAX_TOPIC = 5
)

type RssDataItem struct {
	Title         string `json:"title"`
	PublishedDate string `json:"pubDate"`
	Link          string `json:"link"`
}
type RssResponse struct {
	Status string        `json:"status"`
	Data   []RssDataItem `json:"items"`
}

type GithubActivity struct {
	TypeEvent string `json:"type"`
	Repo      struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Commits []struct {
			Msg        string `json:"message"`
			CommitHash string `json:"sha"`
			URL        string `json:"url"`
			HTMLURL    string `json:"html_url"`
		} `json:"commits"`
	} `json:"payload"`
	CreatedAt string `json:"created_at"`
}

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsJSON([]byte(os.Getenv("FIRESTORE_KEY")))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	iter := client.Collection("gophers").Documents(ctx)
	var usernames []string
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		data := doc.Data()
		usernames = append(usernames, data["username"].(string))
	}

	var rssData []RssDataItem
	for _, username := range usernames {
		res, err := http.Get(fmt.Sprintf(
			"http://api.rss2json.com/v1/api.json?rss_url=https://medium.com/feed/@%v", username))
		if err != nil {
			log.Println("error when call medium API username", username, "error:", err)
			continue
		}
		data, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		rssResponse := RssResponse{}
		err = json.Unmarshal(data, &rssResponse)
		if err != nil {
			log.Println("error unmarshal data", username, "error:", err)
			continue
		}
		rssData = append(rssData, rssResponse.Data...)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rssData), func(i, j int) { rssData[i], rssData[j] = rssData[j], rssData[i] })
	rssData = rssData[:MAX_TOPIC]
	sort.Slice(rssData, func(i, j int) bool { return rssData[i].PublishedDate > rssData[j].PublishedDate })

	// Send telebot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	log.Printf("Authorized on account %s \n", bot.Self.UserName)
	if err != nil {
		log.Println("error connect telegram, error", err)
	} else {
		id, _ := strconv.ParseInt(os.Getenv("TELEGRAM_BOT_CHAT_ID"), 10, 64)
		var sb strings.Builder
		sb.WriteString("Today publications\n")
		for _, rss := range rssData {
			sb.WriteString(rss.Title + " - " + rss.Link + " - " + rss.PublishedDate)
			sb.WriteRune('\n')
		}
		msg := tgbotapi.NewMessage(id, sb.String())
		bot.Send(msg)
	}

	githubActivities, err := getGithubRecentActivity("duyquang6")
	writeReadme(rssData, githubActivities)
}

func getGithubRecentActivity(username string) ([]GithubActivity, error) {
	res, err := http.Get(fmt.Sprintf(
		"https://api.github.com/users/%v/events", username))
	if err != nil {
		log.Println("error when call github API username", username, "error:", err)
		return nil, err
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	githubActivities := []GithubActivity{}
	resp := []GithubActivity{}
	err = json.Unmarshal(data, &githubActivities)
	// For moment, filter out PushEvent only
	count := 0
	for _, val := range githubActivities {
		if val.TypeEvent == "PushEvent" && val.Repo.Name != "duyquang6/duyquang6" {
			val.Payload.Commits[0].HTMLURL = fmt.Sprintf("https://github.com/%v/commit/%v", val.Repo.Name, val.Payload.Commits[0].CommitHash)
			resp = append(resp, val)
			count++
			if count == 5 {
				break
			}
		}
	}
	return resp, err
}

func writeReadme(rssData []RssDataItem, githubActivities []GithubActivity) {
	// http://api.openweathermap.org/data/2.5/weather?id=1566083&appid={OPENWEATHER_APIKEY}
	type TemplateData struct {
		Today            string
		RssData          []RssDataItem
		GithubActivities []GithubActivity
	}
	readmeFile, err := os.Create("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer readmeFile.Close()
	tmpl := template.Must(template.ParseFiles("template.md"))
	err = tmpl.Execute(readmeFile, TemplateData{
		Today:            time.Now().Format("02-Jan-2006"),
		RssData:          rssData,
		GithubActivities: githubActivities,
	})
	if err != nil {
		panic(err)
	}
}
