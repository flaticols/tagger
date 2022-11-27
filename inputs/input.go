package inputs

import (
	"fmt"
	"strconv"

	gha "github.com/sethvargo/go-githubactions"
)

type Inputs struct {
	GitHubToken       string
	PullRequestNumber int
	DefaultTag        string
	TagPattern        string
}

// GetInputs returns the inputs for the action
func GetInputs() (Inputs, error) {
	prNumberInput := gha.GetInput("pr-number")
	prNumber, err := strconv.Atoi(prNumberInput)

	if err != nil {
		return Inputs{}, fmt.Errorf("get PR number failed, %w", err)
	}

	return Inputs{
		GitHubToken:       gha.GetInput("github-token"),
		PullRequestNumber: prNumber,
		DefaultTag:        gha.GetInput("default-tag"),
		TagPattern:        gha.GetInput("tag-pattern"),
	}, nil
}
