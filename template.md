# Welcome to my github 
Today is {{.Today}}\
The language I used recently:\
[![Top Langs](https://github-readme-stats.vercel.app/api/top-langs/?username=duyquang6&layout=compact&hide=html&theme=dark)](https://github.com/anuraghazra/github-readme-stats)\
What I do recently on github:\
![duyquang6's github stats](https://github-readme-stats.vercel.app/api?username=duyquang6&layout=compact&theme=dark&hide=stars,prs,contribs,issues)
{{range .GithubActivities}}
 - Contributing repo {{.Repo.Name}} with [commit]({{(index .Payload.Commits 0).HTMLURL}}) `{{(index .Payload.Commits 0).Msg}}` on  {{.CreatedAt}}
{{end}}
There is some article you might want to read:
{{range .RssData}}
 - [{{.Title}} - {{.PublishedDate}}]({{.Link}})
{{end}}
