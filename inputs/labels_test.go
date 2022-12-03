package inputs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetNewTag_Patch(t *testing.T) {
	tag := "v0.4.3"

	newTag, err := GetNewTag(tag, SemVerUpdate{Patch: true})
	require.NoError(t, err)
	require.Equal(t, "v0.4.4", newTag)
}

func TestGetNewTag_Major(t *testing.T) {
	tag := "v0.4.3"

	newTag, err := GetNewTag(tag, SemVerUpdate{Minor: true, Patch: true})
	require.NoError(t, err)
	require.Equal(t, "v0.5.0", newTag)
}
func TestGetNewTag_Minor(t *testing.T) {
	tag := "v0.4.3"

	newTag, err := GetNewTag(tag, SemVerUpdate{Major: true, Minor: true, Patch: true})
	require.NoError(t, err)
	require.Equal(t, "v1.0.0", newTag)
}

func TestGetNewTag_Custom(t *testing.T) {
	tag := "tag v0.4.3-beta"

	newTag, err := GetNewTag(tag, SemVerUpdate{Major: true, Minor: true, Patch: true})
	require.NoError(t, err)
	require.Equal(t, "v0.4.3-beta", newTag)
}

func TestGetNewTag_Custom1(t *testing.T) {
	tag := "tag:v0.4.3-beta"

	newTag, err := GetNewTag(tag, SemVerUpdate{Major: true, Minor: true, Patch: true})
	require.NoError(t, err)
	require.Equal(t, "v0.4.3-beta", newTag)
}

func TestGetNewTag_Custom2(t *testing.T) {
	tag := "tag-v0.4.3-beta"

	newTag, err := GetNewTag(tag, SemVerUpdate{Major: true, Minor: true, Patch: true})
	require.NoError(t, err)
	require.Equal(t, "v0.4.3-beta", newTag)
}

func TestGetNewTag_Custom3(t *testing.T) {
	tag := "tag=v0.4.3-beta"

	newTag, err := GetNewTag(tag, SemVerUpdate{Major: true, Minor: true, Patch: true})
	require.NoError(t, err)
	require.Equal(t, "v0.4.3-beta", newTag)
}
