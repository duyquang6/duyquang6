# Welcome to my github 
Today is {{.Today}}\
What I do recently on github:
{{range .GithubActivities}}
 - Contributing repo {{.Repo.Name}} with [commit]({{(index .Payload.Commits 0).HTMLURL}}) `{{(index .Payload.Commits 0).Msg}}` on  {{.CreatedAt}}
{{end}}
There is some article you might want to read:
{{range .RssData}}
 - [{{.Title}} - {{.PublishedDate}}]({{.Link}})
{{end}}