package inputs

import (
	"fmt"
	"github.com/google/go-github/v48/github"
	"github.com/hashicorp/go-version"
	"regexp"
	"strings"
)

var (
	semVerPreMetaPattern = regexp.MustCompile(`(?P<pre>pre:([a-zA-z0-9\-\\.]+))|(?P<meta>meta:([a-zA-z0-9\-\\.]+))`)
)

type PullRequestLabels struct {
	Major      bool
	Minor      bool
	Patch      bool
	Prerelease string
	Metadata   string
	Custom     string
}

// GetPRLabels returns a list of labels for a PR
func GetPRLabels(labels []*github.Label) PullRequestLabels {
	sm := PullRequestLabels{Major: false, Minor: false, Patch: false}

	for _, label := range labels {
		name := strings.ToLower(*label.Name)

		switch name {
		case "major":
			sm.Major = true
		case "minor":
			sm.Minor = true
		case "patch":
			sm.Patch = true
		case "pre":
		case "metadata":
		}
	}

	return sm
}

// CreateNewVersion returns a new tag based on the current tag and the PullRequestLabels
func CreateNewVersion(ver string, prLabels PullRequestLabels) (*version.Version, error) {
	semver, err := version.NewSemver(ver)
	if err != nil {
		return nil, fmt.Errorf("invalid version '%s', %w", ver, err)
	}

	segs := semver.Segments()

	if prLabels.Major {
		segs[0]++
		segs[1] = 0
		segs[2] = 0
	} else if prLabels.Minor {
		segs[1]++
		segs[2] = 0
	} else if prLabels.Patch {
		segs[2]++

		return semver, nil
	}

	return version.NewVersion(fmt.Sprintf("%d.%d.%d", segs[0], segs[1], segs[2]))
}
