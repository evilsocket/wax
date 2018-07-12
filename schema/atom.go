package schema

import (
	"reflect"
)

// Atom represents a single element of an input document
// that can be mutated.
type Atom struct {
	// who
	Name string
	// what
	Type reflect.Kind
	// where
	Locator Locator
	// TODO: implement custom mutation constraints ( "but" )

	// this will be updated according to the fitness estimation
	// pipeline, some differential evaluation magic and eventually
	// by the feedback loop
	relevance float32
}
