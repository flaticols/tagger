package inputs

import (
	"github.com/google/go-github/v48/github"
	"regexp"
	"strconv"
	"strings"
)

const (
	MajorLabel = "major"
	MinorLabel = "minor"
	PatchLabel = "patch"
)

type SemVerUpdate struct {
	Major  bool
	Minor  bool
	Patch  bool
	Custom string
}

func GetNextVersionUpdate(labels []*github.Label) SemVerUpdate {
	sm := SemVerUpdate{
		Major:  false,
		Minor:  false,
		Patch:  false,
		Custom: "",
	}

	for _, label := range labels {
		name := strings.ToLower(*label.Name)
		if name == MajorLabel {
			sm.Major = true
		} else if name == MinorLabel {
			sm.Minor = true
		} else if name == PatchLabel {
			sm.Patch = true
		}
	}

	return sm
}

func GetNewTag(tag string, ver SemVerUpdate) (string, error) {
	tagPattern := regexp.MustCompile(`(?P<major>[0-9])\.(?P<minor>[0-9])\.(?P<patch>[0-9])`)

	parts := tagPattern.FindStringSubmatch(tag)

	paramsMap := make(Tag)
	for i, name := range tagPattern.SubexpNames() {
		if i > 0 && i <= len(parts) {
			p, err := strconv.Atoi(parts[i])
			if err != nil {
				return tag, err
			}

			paramsMap[name] = p
		}
	}

	if ver.Major {
		paramsMap["major"]++
		paramsMap["minor"] = 0
		paramsMap["patch"] = 0

		return paramsMap.String(), nil
	}

	if ver.Minor {
		paramsMap["minor"]++
		paramsMap["patch"] = 0
		return paramsMap.String(), nil
	}

	paramsMap["patch"]++
	return paramsMap.String(), nil
}
