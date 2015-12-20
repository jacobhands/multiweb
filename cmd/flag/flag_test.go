package flag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagNames(t *testing.T) {
	flags := map[string]string{
		BaseURL: "baseurl",
		Folder:  "folder",
	}
	for k, v := range flags {
		assert.Equal(t, k, v)
	}
}
