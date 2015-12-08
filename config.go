package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v1"

	"github.com/imdario/mergo"
)

type githubSourceConfig struct {
	Locations map[string]string `yaml:,flow`
}

type config struct {
	Use   []string
	Asana struct {
		Team    string
		Project string
		User    string
		APIKey  string `yaml:"api_key"`
	}
	GithubSource githubSourceConfig `yaml:"github_source"`
	Slack        struct {
		Provider string
	}
}

func contains(s []string, match string) bool {
	for _, a := range s {
		if a == match {
			return true
		}
	}
	return false
}

// Zeroed values in "new" overwritten by those in "preferred"
// New config struct is returned, originals are not modified.
// Useful pattern —
// Given: default, user, and directory configs,
// first := mergeConfig(*directory, *user)
// second := mergeConfig(*first, *default)
func mergeConfig(preferred config, new config) config {
	mergo.Merge(&preferred, new)
	return preferred
}

// Hand it an existing config (perhaps some default values? Just a thought.)
// Also a `home` directory and a `start`, which would typically be the current
// directory. Recursive config search will HALT should it reach that. Searches
// that begin outside the tree under the `home` location will go all the way
// to root (/).
// Also stops looking for a directory config when it finds one, so no nested
// directory configs.

// OK. Now that it's actually working, I hate this entire function.
// TODO: replace this garbage.
func synthesizeConfig(fallback config, home string, current string) (config, error) {
	originalWorkingDirectory, err := os.Getwd()
	if err != nil {
		return fallback, err
	}
	defer os.Chdir(originalWorkingDirectory)
	configFiles := []string{".projecter.yaml", ".projecter"}
	stopAt := []string{home, "/"}
	homeConfig := config{}
	dirConfig := config{}

	// TODO: factor this and the below out into own function. No patience for
	// it now, just want to get something running.
	for _, v := range configFiles {
		path := string(append([]rune(home), filepath.Separator)) + v
		exists := fileExists(path)
		if exists {
			fmt.Println("hit on " + path)
			rawConfig, err := ioutil.ReadFile(path)
			if err != nil {
				return fallback, err
			}
			err = yaml.Unmarshal(rawConfig, &homeConfig)
			if err != nil {
				return fallback, err
			}
			break
		}
	}

	// Gross easily-broken conditional, sorry.
	for current != stopAt[0] && current != stopAt[1] {
		rawConfig := []byte{}
		for _, v := range configFiles {
			path := string(append([]rune(current), filepath.Separator)) + v
			//fmt.Println("Checking path " + path)
			exists := fileExists(path)
			if exists {
				//fmt.Println("hit on " + path)
				rawConfig, err = ioutil.ReadFile(path)
				if err != nil {
					return fallback, err
				}
				err = yaml.Unmarshal(rawConfig, &dirConfig)
				if err != nil {
					return fallback, err
				}
				break
			}
		}
		if len(rawConfig) != 0 {
			break
		}
		if err = os.Chdir(string(append([]rune(current), filepath.Separator)) + "../"); err != nil {
			return fallback, err
		}
		current, err = os.Getwd()
	}

	c := mergeConfig(dirConfig, homeConfig)
	c = mergeConfig(c, fallback)

	return c, nil
}

func getConfig(path string) config {
	config := config{}
	return config
}

func generateRoutes(c config) {

}

// Not propogating up the error because I don't really care why it failed.
// Maybe I should. But I don't.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
