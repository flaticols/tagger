package inputs

import "fmt"

type Tag map[string]int

func (t *Tag) GetMajor() int {
	return (*t)["major"]
}

func (t *Tag) GetMinor() int {
	return (*t)["minor"]
}

func (t *Tag) GetPatch() int {
	return (*t)["patch"]
}

func (t *Tag) String() string {
	return fmt.Sprintf("v%d.%d.%d", t.GetMajor(), t.GetMinor(), t.GetPatch())
}
