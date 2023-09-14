package main

import "fmt"

var discord Service = Service{
	tmpl: mdTmpl,
	buildURL: func(bld Build) string {

		discordID := getenvStrict("PLUGIN_DISCORD_WEBHOOK_ID")
		discordToken := getenvStrict("PLUGIN_DISCORD_WEBHOOK_TOKEN")

		var color string

		switch bld.Status {
		case "SUCCESS":
			color = colorSuccess
		case "FAILURE":
			color = colorFailure
		default:
			color = colorUnknown
		}

		return fmt.Sprintf("discord://%s@%s?color=%s&splitLines=false", discordToken, discordID, color)

	},
}
