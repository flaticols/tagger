package gh

import (
	"context"
	"github.com/google/go-github/v48/github"
)

// CreateRelease creates a new release
func (gh *Client) CreateRelease(tag, owner, repository string, draft, notes bool) (*github.RepositoryRelease, error) {
	rel, _, err := gh.client.Repositories.CreateRelease(context.Background(), owner, repository, &github.RepositoryRelease{
		TagName:              &tag,
		Name:                 &tag,
		Body:                 nil,
		Draft:                &draft,
		Prerelease:           nil,
		GenerateReleaseNotes: &notes,
	})

	return rel, err
}
