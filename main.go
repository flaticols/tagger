package main

import (
	"github.com/flaticols/tagger/inputs"
	gha "github.com/sethvargo/go-githubactions"
	"log"
)

type Opts struct {
	GitHubAction string `command:"github-action" description:"GitHub Action"`
}

func main() {
	params, err := inputs.GetInputs()

	if err != nil {
		log.Fatalln(err.Error())
	}

	gha.Noticef("PR Number: %d, Default Tag: %s", params.PullRequestNumber, params.DefaultTag)
}
