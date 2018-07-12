package schema

// Atom represents a single element of an input document
// that can be mutated.
type Atom struct {
	Name    string  // who
	Type    string  // what
	Locator Locator // where

	// TODO: implement custom mutation constraints ( "but" )

	// this will be updated according to the fitness estimation
	// pipeline, some differential evaluation magic and eventually
	// by the feedback loop
	relevance float32
}
