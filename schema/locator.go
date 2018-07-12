package schema

import (
	"bytes"
	"fmt"
	"regexp"
	"sync"
)

type LocatorType string

// TODO: define and implement more locator types (?)
const (
	ByOffset LocatorType = "offset"
	ByExact  LocatorType = "exact"
	ByRE     LocatorType = "re"
)

// Locator represents a strategy to locate a single
// atom inside a document, by its offset, a XPath
// expression, etc.
type Locator struct {
	sync.Mutex

	Type LocatorType
	Data string

	parsed bool
	offset uint
	size   uint
	re     *regexp.Regexp
}

func (l Locator) byOffset(doc Document) (offset uint, size uint, err error) {
	l.Lock()
	defer l.Unlock()

	// parse offset and size from the Data field
	if !l.parsed {
		if _, err = fmt.Sscanf(l.Data, "%d:%d", &l.offset, &l.size); err != nil {
			return 0, 0, fmt.Errorf("error while parsing locator data '%s' as offset:size: %v", l.Data, err)
		}
		l.parsed = true
	}

	offEnd := l.offset + l.size

	if l.offset >= doc.Size {
		return 0, 0, fmt.Errorf("locator offset %d exceeds document size")
	} else if offEnd >= doc.Size {
		return 0, 0, fmt.Errorf("locator end offset %d exceeds document size")
	}

	return l.offset, l.size, nil
}

func (l Locator) byExact(doc Document) (offset uint, size uint, err error) {
	l.Lock()
	defer l.Unlock()

	// basic stuff
	if !l.parsed {
		l.size = uint(len(l.Data))
		l.parsed = true
	}

	idx := bytes.Index(doc.Data, []byte(l.Data))
	if idx == -1 {
		return 0, 0, fmt.Errorf("locator exact token '%s' could not be found in document", l.Data)
	}
	return uint(idx), l.size, nil
}

func (l Locator) byRegexp(doc Document) (offset uint, size uint, err error) {
	l.Lock()
	defer l.Unlock()

	// basic stuff
	if !l.parsed {
		if l.re, err = regexp.Compile(l.Data); err != nil {
			return 0, 0, fmt.Errorf("could not compile locator expression '%s': %v", l.Data, err)
		}
		l.parsed = true
	}

	m := l.re.FindSubmatchIndex(doc.Data)
	if len(m) < 4 {
		return 0, 0, fmt.Errorf("locator re '%s' could not find a submatch match in document", l.Data)
	}

	from := m[2]
	to := m[3]

	return uint(from), uint(to - from), nil
}

func (l Locator) Find(doc Document) (offset uint, size uint, err error) {
	switch l.Type {
	case ByOffset:
		return l.byOffset(doc)
	case ByExact:
		return l.byExact(doc)
	case ByRE:
		return l.byRegexp(doc)
	}

	return 0, 0, fmt.Errorf("unhandled schema.Locator type '%s'", l.Type)
}
