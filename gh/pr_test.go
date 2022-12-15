package gh

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNormalizeRepositoryName(t *testing.T) {
	repository := "flaticols/tagger"
	owner := "flaticols"

	repository = strings.Replace(repository, fmt.Sprintf("%s/", owner), "", -1)

	require.Equal(t, "tagger", repository)
}
