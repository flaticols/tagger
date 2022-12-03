package action

import (
	"context"
	"github.com/google/go-github/v48/github"
)

func (gh *Client) CreateRelease(tag, owner, repository string) (*github.RepositoryRelease, error) {
	rel, _, err := gh.client.Repositories.CreateRelease(context.Background(), owner, repository, &github.RepositoryRelease{
		TagName:              &tag,
		Name:                 &tag,
		Body:                 nil,
		Draft:                false,
		Prerelease:           nil,
		GenerateReleaseNotes: true,
	})

	return rel, err
}
