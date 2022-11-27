package main

import (
	gha "action.com/sethvargo/go-githubactions"
	"github.com/flaticols/tagger/action"
	"github.com/flaticols/tagger/inputs"
	"log"
)

type Opts struct {
	GitHubAction string `command:"action-action" description:"GitHub Action"`
}

func main() {
	gitHubContext, err := gha.Context()
	if err != nil {
		log.Fatalf("get GitHub context failed, %s", err.Error())
	}

	owner := gitHubContext.RepositoryOwner
	repo := gitHubContext.Repository

	params, err := inputs.GetInputs()

	if err != nil {
		log.Fatalf("get inputs failed, %s", err.Error())
	}

	client, err := action.NewGitHubClient(params.GitHubToken)
	if err != nil {
		log.Fatalf("failed to create GitHub client, %s", err.Error())
	}

	labels, err := client.GetPullRequestLabels(params.PullRequestNumber, owner, repo)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ver := inputs.GetNextVersionUpdate(labels)
	tag, err := client.GetLatestTag(owner, repo)
	if err != nil {
		log.Fatalf("get latest tag failed, %s", err.Error())
	}

	newTag, err := inputs.GetNewTag(tag, ver)
	if err != nil {
		log.Printf("create new tag failed, %s", err.Error())
	}

	gha.Noticef("PR Number: %d, Default Tag: %s", params.PullRequestNumber, params.DefaultTag)
}
