package inputs

import (
	"fmt"
	"strconv"
	"strings"

	gha "github.com/sethvargo/go-githubactions"
)

type Inputs struct {
	GitHubToken       string
	Owner             string
	Repository        string
	PullRequestNumber int
	DefaultTag        string
	TagPrefix         string
}

// GetInputs returns the inputs for the gh
func GetInputs() (Inputs, error) {
	prNumberInput := gha.GetInput("pr-number")
	prNumber, err := strconv.Atoi(prNumberInput)

	if err != nil {
		return Inputs{}, fmt.Errorf("get PR number failed, %w", err)
	}

	ghaCtx, err := gha.Context()
	if err != nil {
		return Inputs{}, fmt.Errorf("get GitHub context failed, %w", err)
	}

	return Inputs{
		GitHubToken:       gha.GetInput("github-token"),
		Repository:        strings.Split(ghaCtx.Repository, "/")[1],
		Owner:             ghaCtx.RepositoryOwner,
		PullRequestNumber: prNumber,
		DefaultTag:        gha.GetInput("default-tag"),
		TagPrefix:         gha.GetInput("prefix"),
	}, nil
}
