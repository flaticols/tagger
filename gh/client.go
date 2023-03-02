package gh

import (
	"context"
	"fmt"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

type Client struct {
	client *github.Client
}

// NewGitHubClient returns a new GitHub client
func NewGitHubClient(token string) (*Client, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &Client{client: client}, nil
}

// GetLatestTag returns the latest tag for a given repository
func (gh *Client) GetLatestTag(owner, repository string, defaultTag string) (string, error) {
	tags, _, err := gh.client.Repositories.ListTags(context.Background(), owner, repository, &github.ListOptions{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		return "", fmt.Errorf("get latest tag failed, %w", err)
	}

	if len(tags) == 0 {
		return defaultTag, nil
	} else {
		return tags[0].GetName(), nil
	}
}
