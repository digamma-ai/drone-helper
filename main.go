package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkErrorFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	cacheFlags := flag.NewFlagSet("cache", flag.ExitOnError)
	cacheDepsArg := cacheFlags.String("deps", "", "List of files on which cache depends (space-separated)")

	notifyFlags := flag.NewFlagSet("notify", flag.ExitOnError)
	notifyDiscordArg := notifyFlags.Bool("discord", true, `Send notification to Discord. (requires: discord_webhook_token, discord_webhook_id)`)

	if len(os.Args) < 2 {
		log.Fatalln("expected 'cache' or 'notify' subcommands")
	}

	switch os.Args[1] {

	case "cache":
		cacheFlags.Parse(os.Args[2:])
		cacheDeps := strings.Fields(*cacheDepsArg)
		if len(cacheDeps) <= 0 {
			log.Fatalln("cannot indentify cache: no dependencies given (--deps)")
		}
		rebuildCache(cacheDeps)
	case "notify":
		notifyFlags.Parse(os.Args[2:])
		switch {
		case *notifyDiscordArg:
			discord.notify()
		}
	default:
		log.Fatalln(fmt.Sprintf("unexpected subcommand: %s", os.Args[1]))
	}
}
