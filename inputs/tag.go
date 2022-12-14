package inputs

import (
	"fmt"
	"strings"
)

const (
	TagMajorName string = "major"
	TagMinorName string = "minor"
	TagPatchName string = "patch"
)

type Tag struct {
	majorTagName  string
	mainorTagName string
	patchTagName  string

	major  int
	minor  int
	patch  int
	custom string
}

// Set sets the custom version
func (t *Tag) Set(name string, value int) {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case TagMajorName:
		t.SetMajor(value)
	case TagMinorName:
		t.SetMinor(value)
	case TagPatchName:
		t.SetPatch(value)
	}
}

// SetMajor sets the major version
func (t *Tag) SetMajor(val int) {
	t.major = val
}

// SetMinor sets the minor version
func (t *Tag) SetMinor(val int) {
	t.minor = val
}

// SetPatch sets the patch version
func (t *Tag) SetPatch(val int) {
	t.patch = val
}

// GetMajor sets the major version
func (t *Tag) GetMajor() int {
	return t.major
}

// GetMinor sets the minor version
func (t *Tag) GetMinor() int {
	return t.minor
}

// GetPatch sets the patch version
func (t *Tag) GetPatch() int {
	return t.patch
}

func (t *Tag) String() string {
	return fmt.Sprintf("%d.%d.%d", t.major, t.minor, t.patch)
}

func (t *Tag) StringWithPrefix(prefix string) string {
	return fmt.Sprintf("%s%d.%d.%d", prefix, t.major, t.minor, t.patch)
}
