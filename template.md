# Welcome to my github 
Today is {{.Today}}\
The programming languages I used recently:\
<img src="https://wakatime.com/share/@duyquang6/fbe267a6-a29b-4a1a-b769-c566a361c376.svg" width="600">
What I do recently on github:\
![duyquang6's github stats](https://github-readme-stats.vercel.app/api?username=duyquang6&layout=compact&hide=stars,prs,contribs,issues)
{{range .GithubActivities}}
 - Contributing repo {{.Repo.Name}} with [commit]({{(index .Payload.Commits 0).HTMLURL}}) `{{(index .Payload.Commits 0).Msg}}` on  {{.CreatedAt}}
{{end}}
There is some article you might want to read:
{{range .RssData}}
 - [{{.Title}} - {{.PublishedDate}}]({{.Link}})
{{end}}
