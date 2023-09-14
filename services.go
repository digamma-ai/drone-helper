package main

import (
	"fmt"
	"text/template"
)

func color(bld Build) string {
	switch bld.Status {
	case "SUCCESS":
		return "4BB543"
	case "FAILURE":
		return "FC100D"
	default:
		return "808080"
	}
}

var discordTmpl = template.Must(template.New("discord").Parse(
	`**Build [#{{.BuildNumber}}]({{.BuildLink}})**
**of [{{.Repo}}]({{.RepoLink}}):{{.Branch}}@[{{slice .CommitAfter 0 7}}]({{.CommitLink}})**
**by {{.Author}}**
**{{.Status}}**
`))

func discordURL(bld Build) string {
	id := getenvStrict("PLUGIN_DISCORD_WEBHOOK_ID")
	token := getenvStrict("PLUGIN_DISCORD_WEBHOOK_TOKEN")
	return fmt.Sprintf("discord://%s@%s?color=0x%s&splitLines=false", token, id, color(bld))
}

// https://containrrr.dev/shoutrrr/v0.8/services/overview/
var discord Service = Service{
	tmpl:     discordTmpl,
	buildURL: discordURL,
}

var slackTmpl = template.Must(template.New("slack").Parse(
	`*Build <{{.BuildLink}}|#{{.BuildNumber}}>*
*of <{{.RepoLink}}|{{.Repo}}>:{{.Branch}}@<{{.CommitLink}}|{{slice .CommitAfter 0 7}}>*
*by {{.Author}}*
*{{.Status}}*
`))

func slackURL(bld Build) string {
	token := getenvStrict("PLUGIN_SLACK_WEBHOOK_TOKEN")
	return fmt.Sprintf("slack://hook:%s@webhook?color=%%23%s", token, color(bld))
}

var slack Service = Service{
	tmpl:     slackTmpl,
	buildURL: slackURL,
}
