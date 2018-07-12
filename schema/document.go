package schema

import (
	"io/ioutil"
)

type Document struct {
	Path string
	Data []byte // aren't we all just a pointless array of bytes in this entropic universe after all?
	Size uint
	Gen  uint // generation number
}

func LoadDocument(path string) (doc Document, err error) {
	doc = Document{Path: path}
	if doc.Data, err = ioutil.ReadFile(path); err == nil {
		doc.Size = uint(len(doc.Data))
	}
	return doc, err
}
