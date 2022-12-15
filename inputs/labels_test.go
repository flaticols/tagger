package inputs

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetNewTag_Patch(t *testing.T) {
	tag := "0.4.3"

	newTag, err := CreateNewVersion(tag, PullRequestLabels{Patch: true})
	require.NoError(t, err)
	require.Equal(t, "0.4.4", newTag.String())
}

func TestGetNewTag_Major(t *testing.T) {
	tag := "0.4.3"

	newTag, err := CreateNewVersion(tag, PullRequestLabels{Minor: true, Patch: true})
	require.NoError(t, err)
	require.Equal(t, "0.5.0", newTag.String())
}

func TestGetNewTag_Minor(t *testing.T) {
	tag := "0.4.3"

	newTag, err := CreateNewVersion(tag, PullRequestLabels{Major: true, Minor: true, Patch: true})
	require.NoError(t, err)
	require.Equal(t, "1.0.0", newTag.String())
}

func TestParsePrereleaseAndMetadata(t *testing.T) {
	labelValue := `A software release is often a culmination of hard work and dedication from the development team.
pre:test.452 It represents a major milestone in the evolution of the software, bringing new features and improvements that enhance the user experience.
As the software is unleashed into the world, it has the potential to make a meaningful impact on the lives of those who use it (meta:build.123).
With each release, the software continues to grow and evolve, constantly pushing the boundaries of what is possible.
The release of this software is a moment to be celebrated, as it brings us one step closer to a better, more connected world with
pre:test.452 foo`

	parts := semVerPreMetaPattern.FindStringSubmatch(labelValue)
	fmt.Println(parts)
	t.Logf("Meta: %s", parts[4])
}

func TestCreateNewVersion_Custom(t *testing.T) {
	val := "0.0.1-test.434"

	ver, err := version.NewSemver(val)
	require.NoError(t, err, "failed to create version")

	t.Logf("Core: %s", ver.Core().String())
	t.Logf("Prerelease: %s", ver.Prerelease())
	t.Logf("Metadata: %s", ver.Metadata())

	prerelease := ver.Prerelease()
	if prerelease != "" {
		prerelease = fmt.Sprintf("-%s", prerelease)
	}

	metadata := ver.Metadata()
	if metadata != "" {
		metadata = fmt.Sprintf("+%s", metadata)
	}

	segs := ver.Segments()

	newVer, err := version.NewSemver(fmt.Sprintf("%d.%d.%d%s%s", segs[0], segs[1], segs[2], prerelease, metadata))
	require.NoError(t, err, "failed to create new version")

	t.Log(newVer.String())
}
