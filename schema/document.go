package schema

import (
	"io/ioutil"
)

// aren't we all just a pointless array of bytes in this entropic universe after all?
type Document []byte

func LoadDocument(path string) (Document, error) {
	data, err := ioutil.ReadFile(path)
	return Document(data), err
}
