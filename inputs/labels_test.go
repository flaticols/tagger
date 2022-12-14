package inputs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetNewTag_Patch(t *testing.T) {
	tag := "0.4.3"

	newTag, err := GetNewTag(tag, PullRequestLabels{Patch: true})
	require.NoError(t, err)
	require.Equal(t, "0.4.4", newTag.String())
}

func TestGetNewTag_Major(t *testing.T) {
	tag := "0.4.3"

	newTag, err := GetNewTag(tag, PullRequestLabels{Minor: true, Patch: true})
	require.NoError(t, err)
	require.Equal(t, "0.5.0", newTag.String())
}

func TestGetNewTag_Minor(t *testing.T) {
	tag := "0.4.3"

	newTag, err := GetNewTag(tag, PullRequestLabels{Major: true, Minor: true, Patch: true})
	require.NoError(t, err)
	require.Equal(t, "1.0.0", newTag.String())
}
