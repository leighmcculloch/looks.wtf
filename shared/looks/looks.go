package looks

var tags = func() []string {
	tagsYaml := files["tags.yml"]
	if len(tagsYaml) == 0 {
		panic("tags.yml file missing")
	}

	return loadTags(tagsYaml)
}()

var looksByTags = func() map[string][]Look {
	looksYaml := files["looks.yml"]
	if len(looksYaml) == 0 {
		panic("looks.yml file missing")
	}
	return loadLooks(looksYaml)
}()

func Tags() []string {
	return tags
}

func LooksWithTag(tag string) []Look {
	return looksByTags[tag]
}
