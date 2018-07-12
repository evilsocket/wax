package schema

import (
	"encoding/json"
	"io/ioutil"
)

// Descriptor describes the structure of an input document
// in terms of atoms, their types, how to locate them, and
// optional mutation constraints.
type Descriptor struct {
	Name  string // name of this schema descriptor (pdf, zip, websocket, whatever)
	Atoms []Atom // what the document with this schema is made of
}

func LoadDescriptor(path string) (d Descriptor, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &d)
	return
}

func (d Descriptor) Prepare(doc Document) error {
	for _, atom := range d.Atoms {
		// check if the locator can isolate this mf
		if off, size, err := atom.Locator.Find(doc); err != nil {
			return err
		} else if err = atom.Prepare(doc.Data[off : off+size]); err != nil {
			return err
		}
	}
	return nil
}
