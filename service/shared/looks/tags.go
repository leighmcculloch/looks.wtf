package looks

import (
	"io"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func ParseTags(r io.Reader) []string {
	var tags []string
	dec := yaml.NewDecoder(r)
	err := dec.Decode(&tags)
	if err != nil {
		log.Fatal("Error unmarshaling tags yaml:", err)
	}
	return tags
}
