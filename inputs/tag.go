package inputs

import "fmt"

type TagType string

const (
	TagMajor TagType = "major"
	TagMinor TagType = "minor"
	TagPatch TagType = "patch"
)

type Tag struct {
	major  int
	minor  int
	patch  int
	custom string
}

func (t *Tag) Set(key TagType, value interface{}) {
	switch key {
	case "major":
		t.major = value.(int)
	case "minor":
		t.minor = value.(int)
	case "patch":
		t.patch = value.(int)
	case "custom":
		t.custom = value.(string)
	}
}

func (t *Tag) GetMajor() int {
	return t.major
}

func (t *Tag) GetMinor() int {
	return t.minor
}

func (t *Tag) GetPatch() int {
	return t.patch
}

func (t *Tag) GetCustom() string {
	return t.custom
}

func (t *Tag) String() string {
	if t.GetCustom() != "" {
		return t.GetCustom()
	}

	return fmt.Sprintf("v%d.%d.%d", t.GetMajor(), t.GetMinor(), t.GetPatch())
}
