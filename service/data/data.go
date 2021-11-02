package data

import (
	"bytes"
	_ "embed"
)

//go:embed looks.yml
var looks []byte

var Looks = ParseLooks(bytes.NewReader(looks))

//go:embed tags.yml
var tags []byte

var Tags = ParseTags(bytes.NewReader(tags))
