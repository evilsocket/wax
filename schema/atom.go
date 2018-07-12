package schema

import (
	"sort"
)

type AtomType string

// TODO: define and implement more atom types (?)
const (
	Buffer  AtomType = "buffer"
	String  AtomType = "string"
	Integer AtomType = "int"
	Float   AtomType = "float"
	Bool    AtomType = "bool"
)

// Atom represents a single element of an input document
// that can be mutated.
type Atom struct {
	Name    string   // who
	Type    AtomType // what
	Locator Locator  // where

	// TODO: implement custom mutation constraints ( "but" )

	// this will be updated according to the fitness estimation
	// pipeline, some differential evaluation magic and eventually
	// by the feedback loop
	relevance  float32
	dictionary []byte // list of unique bytes for non scalar types
}

func (a Atom) IsScalar() bool {
	return a.Type != Buffer && a.Type != String
}

func (a Atom) Prepare(raw []byte) error {
	if !a.IsScalar() {
		tmp := make(map[byte]bool)
		for _, b := range raw {
			tmp[b] = true
		}

		a.dictionary = make([]byte, 0)
		for b, _ := range tmp {
			a.dictionary = append(a.dictionary, b)
		}

		sort.Sort(ByteSlice(a.dictionary))
	}

	return nil
}
