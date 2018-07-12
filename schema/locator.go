package schema

type LocatorType uint

// TODO: define and implement more locator types
const (
	ByOffset LocatorType = iota
	ByPrefix
	BySuffix
)

// Locator represents a strategy to locate a single
// atom inside a document, by its offset, a XPath
// expression, etc.
type Locator struct {
	Type LocatorType
	Data string
}
