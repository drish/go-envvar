package envvar

import (
	"os"
	"testing"

	"github.com/tj/assert"
)

func TestLoad_valid(t *testing.T) {
	assert.Equal(t, os.Getenv("PORT"), "")
	Load("", "")
	assert.Equal(t, os.Getenv("PORT"), "5006")
}
