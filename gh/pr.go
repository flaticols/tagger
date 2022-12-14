package gh

import (
	"context"
	"fmt"
	"github.com/google/go-github/v48/github"
)

// GetPullRequestByNumber returns the PR for a given PR number
func (gh *Client) GetPullRequestByNumber(prNumber int, owner, repository string) (*github.PullRequest, error) {
	pr, _, err := gh.client.PullRequests.Get(context.Background(), owner, repository, prNumber)

	if err != nil {
		return nil, fmt.Errorf("get PR '%d' failed, %w", prNumber, err)
	}

	return pr, nil
}

// GetPullRequestLabels returns the labels for a given PR
func (gh *Client) GetPullRequestLabels(prNumber int, owner, repository string) ([]*github.Label, error) {
	pr, err := gh.GetPullRequestByNumber(prNumber, owner, repository)

	if err != nil {
	}

	return pr.Labels, nil
}
