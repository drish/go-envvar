package envvar

import (
	"os"
	"testing"

	"github.com/tj/assert"
)

func TestLoad_panics_when_invalid_path(t *testing.T) {
	assert.Panics(t, func() {
		Load("", "")
	}, "A config file is required")
}

func TestLoad_valid_config(t *testing.T) {
	assert.Equal(t, os.Getenv("PORT"), "")
	Load("./test.yml", "local")
	assert.Equal(t, os.Getenv("PORT"), "5006")
}
