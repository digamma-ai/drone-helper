package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/containrrr/shoutrrr"
)

type Service struct {
	tmpl     *template.Template
	buildURL func(Build) string
}

func (srv Service) notify() {

	bld := getBuildInfo()

	var buf bytes.Buffer
	srv.tmpl.Execute(&buf, bld)
	message := buf.String()

	url := srv.buildURL(bld)

	fmt.Printf("Sending message:\n%s\n", message)
	err := shoutrrr.Send(url, message)
	checkErrorFatal(err)
	fmt.Println("Message sent.")

}
