package inputs

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v48/github"
)

type PullRequestLabels struct {
	Major  bool
	Minor  bool
	Patch  bool
	Custom string

	Tag *Tag
}

// GetPRLabels returns a list of labels for a PR
func GetPRLabels(labels []*github.Label) PullRequestLabels {
	sm := PullRequestLabels{Major: false, Minor: false, Patch: false}

	for _, label := range labels {
		name := strings.ToLower(*label.Name)
		if name == TagMajorName {
			sm.Major = true
		} else if name == TagMinorName {
			sm.Minor = true
		} else if name == TagPatchName {
			sm.Patch = true
		}
	}

	return sm
}

// GetNewTag returns a new tag based on the current tag and the PullRequestLabels
func GetNewTag(currentVersion string, prLabels PullRequestLabels) (Tag, error) {
	tagPattern := regexp.MustCompile(`(?P<major>[0-9]+)\.(?P<minor>[0-9]+)\.(?P<patch>[0-9]+)`)
	parts := tagPattern.FindStringSubmatch(currentVersion)

	tag := Tag{}

	for i, name := range tagPattern.SubexpNames() {
		if i > 0 && i <= len(parts) {
			p, err := strconv.Atoi(parts[i])
			if err != nil {
				return tag, err
			}

			tag.Set(name, p)
		}
	}

	if prLabels.Major {
		tag.SetMajor(tag.GetMajor() + 1)
		tag.SetMinor(0)
		tag.SetPatch(0)
		return tag, nil
	}

	if prLabels.Minor {
		tag.SetMinor(tag.GetMinor() + 1)
		tag.SetPatch(0)
		return tag, nil
	}

	tag.SetPatch(tag.GetPatch() + 1)

	return tag, nil
}
