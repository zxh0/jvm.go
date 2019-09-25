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

func (entry *CompositeEntry) addEntry(subEntry Entry) {
	entry.entries = append(entry.entries, subEntry)
}

func (entry *CompositeEntry) readClass(className string) (Entry, []byte, error) {
	for _, subEntry := range entry.entries {
		subEntry, data, err := subEntry.readClass(className)
		if err == nil {
			return subEntry, data, nil
		}
	}

	return entry, nil, errors.New("class not found: " + className)
}

func (entry *CompositeEntry) String() string {
	strs := make([]string, len(entry.entries))

	for i, subEntry := range entry.entries {
		strs[i] = subEntry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
