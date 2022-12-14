package main

import (
	"github.com/flaticols/tagger/commands"
	"github.com/flaticols/tagger/gh"
	"github.com/flaticols/tagger/inputs"
	gha "github.com/sethvargo/go-githubactions"
	"github.com/spf13/cobra"
	"log"
)

func main_old() {
	params, err := inputs.GetInputs()

	if err != nil {
		log.Fatalf("get inputs failed, %s", err.Error())
	}

	client, err := gh.NewGitHubClient(params.GitHubToken)
	if err != nil {
		log.Fatalf("failed to create GitHub client, %s", err.Error())
	}

	labels, err := client.GetPullRequestLabels(params.PullRequestNumber, params.Owner, params.Repository)
	if err != nil {
		log.Fatalln(err.Error())
	}

	prLabels := inputs.GetPRLabels(labels)
	latestTag, err := client.GetLatestTag(params.Owner, params.Repository)
	if err != nil {
		log.Fatalf("get latest tag failed, %s", err.Error())
	}

	newTag, err := inputs.GetNewTag(latestTag, prLabels)
	if err != nil {
		log.Printf("create new latestTag failed, %s", err.Error())
	}

	gha.Noticef("New latestTag: %s", newTag)
	gha.Noticef("PR Number: %d, Default Tag: %s", params.PullRequestNumber, params.DefaultTag)
}

func main() {
	root := cobra.Command{}

	root.AddCommand(commands.DoCommand())

	_, err := root.ExecuteC()
	if err != nil {
		return
	}
}
