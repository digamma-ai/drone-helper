package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/containrrr/shoutrrr"
)

const colorSuccess = "0x4BB543"
const colorFailure = "0xFC100D"
const colorUnknown = "0x808080"

// template assuming basic markdown support
// (plainer and richer options available, depending on service)
var mdTmpl = template.Must(template.New("markdown-message").Parse(
	`**Build [#{{.BuildNumber}}]({{.BuildLink}})**
**of [{{.Repo}}]({{.RepoLink}}):{{.Branch}}@[{{slice .CommitAfter 0 7}}]({{.CommitLink}})**
**by {{.Author}}**
**{{.Status}}**
`))

func notifyDiscord() {

	bld := getBuildInfo()

	discordID := getenvStrict("DISCORD_WEBHOOK_ID")
	discordToken := getenvStrict("DISCORD_WEBHOOK_TOKEN")

	var color string

	switch bld.Status {
	case "SUCCESS":
		color = colorSuccess
	case "FAILURE":
		color = colorFailure
	default:
		color = colorUnknown
	}

	var buf bytes.Buffer
	mdTmpl.Execute(&buf, bld)
	message := buf.String()

	discordURL := fmt.Sprintf("discord://%s@%s?color=%s&splitLines=false", discordToken, discordID, color)

	fmt.Printf("Sending message to Discord:\n%s\n", message)
	err := shoutrrr.Send(discordURL, message)
	checkErrorFatal(err)
	fmt.Println("Message sent.")

}
