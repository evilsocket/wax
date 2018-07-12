package schema

import (
	"bytes"
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

	dictionary []byte // list of unique bytes for non scalar types
	hasSpaces  bool
	hasSymbols bool
	hasDigits  bool
	hasLetters bool
	hasBinary  bool
	// this will be updated according to the fitness estimation
	// pipeline, some differential evaluation magic and eventually
	// by the feedback loop
	relevance float32
}

func (a Atom) IsScalar() bool {
	return a.Type != Buffer && a.Type != String
}

func (a Atom) Prepare(raw []byte) error {
	if !a.IsScalar() {
		// build a dictionary of the unique bytes
		tmp := make(map[byte]bool)
		for _, b := range raw {
			tmp[b] = true
		}

		a.dictionary = make([]byte, 0)
		for b, _ := range tmp {
			a.dictionary = append(a.dictionary, b)
		}

		sort.Sort(ByteSlice(a.dictionary))

		// check which byte classes are present in the dictionary
		for _, b := range a.dictionary {
			if bytes.IndexByte([]byte(" \t\n\v\f\r"), b) != -1 {
				a.hasSpaces = true
			} else if (b >= 0x21 && b <= 0x2f) || (b >= 0x3a && b <= 0x40) || (b >= 0x5b && b <= 0x60) || (b >= 0x7b && b <= 0x7e) {
				a.hasSymbols = true
			} else if bytes.IndexByte([]byte("1234567890"), b) != -1 {
				a.hasDigits = true
			} else if (b >= 0x41 && b <= 0x5a) || (b >= 0x61 && b <= 0x7a) {
				a.hasLetters = true
			} else {
				a.hasBinary = true
			}
		}
	}

	return nil
}
