package main

import (
	"log"
	"os"
	"strings"
)

type Build struct {
	BuildNumber  string
	BuildEvent   string
	BuildLink    string
	Repo         string
	RepoLink     string
	Branch       string
	CommitBefore string
	CommitAfter  string
	CommitLink   string
	Author       string
	Status       string
}

func getenvStrict(key string) string {
	res := os.Getenv(key)
	if res == "" {
		log.Fatalf("%s: required variable not defined", key)
	}
	return res
}

func getBuildInfo() Build {

	return Build{
		BuildNumber:  getenvStrict("DRONE_BUILD_NUMBER"),
		BuildEvent:   getenvStrict("DRONE_BUILD_EVENT"),
		BuildLink:    getenvStrict("DRONE_BUILD_LINK"),
		Repo:         getenvStrict("DRONE_REPO"),
		RepoLink:     getenvStrict("DRONE_REPO_LINK"),
		Branch:       getenvStrict("DRONE_SOURCE_BRANCH"),
		CommitBefore: getenvStrict("DRONE_COMMIT_BEFORE"),
		CommitAfter:  getenvStrict("DRONE_COMMIT_AFTER"),
		CommitLink:   getenvStrict("DRONE_COMMIT_LINK"),
		Author:       getenvStrict("DRONE_COMMIT_AUTHOR"),
		Status:       strings.ToUpper(getenvStrict("DRONE_BUILD_STATUS")),
	}

}
