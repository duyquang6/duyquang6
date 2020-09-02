# Welcome to my github 
Today is {{.Today}}\
There is some article you might want to read:
{{range .RssData}}
 - [{{.Title}} - {{.PublishedDate}}]({{.Link}})
{{end}}