package main

import (
	"io/ioutil"
	"log"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type look struct {
	Plain string `yaml:"plain"`
	Tags  string `yaml:"tags"`
}

func loadLooks(file string) map[string][]look {
	looksYaml, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error loading looks yaml:", err)
	}
	var looks []look
	err = yaml.Unmarshal(looksYaml, &looks)
	if err != nil {
		log.Fatal("Error unmarshaling looks yaml:", err)
	}

	var looksByTags = make(map[string][]look)
	for _, l := range looks {
		tags := strings.Split(l.Tags, " ")
		for _, t := range tags {
			if t == "" {
				continue
			}
			looksByTag, ok := looksByTags[t]
			if !ok {
				looksByTag = []look{}
			}
			looksByTags[t] = append(looksByTag, l)
		}
	}
	return looksByTags
}
