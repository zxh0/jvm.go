package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// :(linux/unix) or ;(windows)
const _pathListSeparator = string(os.PathListSeparator)

type CompositeEntry struct {
	entries []Entry
}

func newCompositeEntry(pathList string) *CompositeEntry {
	compoundEntry := &CompositeEntry{}

	for _, path := range strings.Split(pathList, _pathListSeparator) {
		if absPath, err := filepath.Abs(path); err == nil {
			entry := parseEntry(absPath)
			compoundEntry.addEntry(entry)
		} else {
			// todo
		}
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

	// todo
	return nil, nil, classNotFoundErr
}

func (self *CompositeEntry) String() string {
	strs := make([]string, len(self.entries))

	for i, entry := range self.entries {
		strs[i] = entry.String()
	}

	return strings.Join(strs, _pathListSeparator)
}
