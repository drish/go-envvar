// Package envvar provides an opnionated way of handling environment variables as yaml files
//
// the yaml config file consist of the following structure:
//
// required:
//   - FOO
//   - BAR
// defaults:
//   BIZ: "set if unset"
// dev:
//   VAR: SOP
// staging:
// prod:
package envvar

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var DEFAULTS_KEY = "defaults"
var REQUIRED_KEY = "required"

// Load configuration file from path, exits if can't load
func Load(path, env string) {
	var c interface{}

	if env == "" {
		env = "development"
	}

	f, err := ioutil.ReadFile("./example.yml")
	if err != nil {
		log.Fatal(err)
	}

	yaml.Unmarshal(f, &c)
	parse(c)
}

// parse checks for the presence of env vars and sets them when necessary
func parse(c interface{}) {
	var requireds []string
	defaults := make(map[string]string)

	for k, v := range c.(map[interface{}]interface{}) {

		// parses required env vars
		if k.(string) == REQUIRED_KEY {
			for _, j := range v.([]interface{}) {
				requireds = append(requireds, j.(string))
			}
		}

		// parses default env vars
		if k.(string) == DEFAULTS_KEY {
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
			log.Fatalf("required env var %s is not set", r)
		}
	}
}
