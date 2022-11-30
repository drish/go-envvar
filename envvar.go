// Package envvar provides an opnionated way of handling environment variables as yaml files
// see https://github.com/drish/go-envvar/blob/master/example.yml for an example config
package envvar

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const defaultsKey = "defaults"
const requiredKey = "required"

// defaultEnvKey is the key that envvar defaults when not supplied by the user
const defaultEnvKey = "local"

// Load configuration file from path, exits if can't load
func Load(path, env string) {
	var c interface{}

	if path == "" {
		panic("A config file is required")
	}

	if env == "" {
		env = defaultEnvKey
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(f, &c)
	Parse(c)
}

// Parse checks for the presence of env vars and sets them when necessary
// TODO: load current env
func Parse(c interface{}) {
	var requireds []string
	defaults := make(map[string]string)

	for k, v := range c.(map[interface{}]interface{}) {

		// parses required env vars
		if k.(string) == requiredKey {
			for _, j := range v.([]interface{}) {
				requireds = append(requireds, j.(string))
			}
		}

		// parses default env vars
		if k.(string) == defaultsKey {
			for i, j := range v.(map[interface{}]interface{}) {
				defaults[i.(string)] = j.(string)
			}
		}
	}

	// set all defaults
	for dk, dv := range defaults {
		if os.Getenv(dk) == "" {
			os.Setenv(dk, dv)
		}
	}

	// check for requireds and exit if any is not set
	for _, r := range requireds {
		if os.Getenv(r) == "" {
			panic("required env var %s is not set")
		}
	}
}
