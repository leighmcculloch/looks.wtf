package slackcommands

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func loadTags(file string) []string {
	tagsYaml, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error loading tags yaml:", err)
	}
	var tags []string
	err = yaml.Unmarshal(tagsYaml, &tags)
	if err != nil {
		log.Fatal("Error unmarshaling looks yaml:", err)
	}
	return tags
}
