package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry struct {
	entries []Entry
}

func newCompositeEntry(pathList string) *CompositeEntry {
	compoundEntry := &CompositeEntry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compoundEntry.addEntry(entry)
	}

	return compoundEntry
}

func (self *CompositeEntry) addEntry(entry Entry) {
	self.entries = append(self.entries, entry)
}

func (self *CompositeEntry) readClass(className string) (Entry, []byte, error) {
	for _, entry := range self.entries {
		entry, data, err := entry.readClass(className)
		if err == nil {
			return entry, data, nil
		}
	}

	return self, nil, errors.New("class not found: " + className)
}

func (self *CompositeEntry) String() string {
	strs := make([]string, len(self.entries))

	for i, entry := range self.entries {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
