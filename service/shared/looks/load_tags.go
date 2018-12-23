package looks

import (
	"log"

	yaml "gopkg.in/yaml.v2"
)

func loadTags(tagsYaml []byte) []string {
	var tags []string
	err := yaml.Unmarshal(tagsYaml, &tags)
	if err != nil {
		log.Fatal("Error unmarshaling tags yaml:", err)
	}
	return tags
}
