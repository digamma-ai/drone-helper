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

type Service struct {
	tmpl     *template.Template
	buildURL func(Build) string
}

func (srv Service) notify() {

	bld := getBuildInfo()

	var buf bytes.Buffer
	mdTmpl.Execute(&buf, bld)
	message := buf.String()

	url := srv.buildURL(bld)

	fmt.Printf("Sending message:\n%s\n", message)
	err := shoutrrr.Send(url, message)
	checkErrorFatal(err)
	fmt.Println("Message sent.")

}
