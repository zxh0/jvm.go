package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// :(linux/unix) or ;(windows)
const _pathListSeparator = string(os.PathListSeparator)

type CompoundEntry struct {
	entries []Entry
}

func newCompoundEntry(pathList string) *CompoundEntry {
	compoundEntry := &CompoundEntry{}

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

func (self *CompoundEntry) addEntry(entry Entry) {
	self.entries = append(self.entries, entry)
}

func (self *CompoundEntry) readClassData(className string) (Entry, []byte, error) {
	for _, entry := range self.entries {
		entry, data, err := entry.readClassData(className)
		if err == nil {
			return entry, data, nil
		}
	}

	// todo
	return nil, nil, classNotFoundErr
}

func (self *CompoundEntry) String() string {
	strs := make([]string, len(self.entries))

	for i, entry := range self.entries {
		strs[i] = entry.String()
	}

	return strings.Join(strs, _pathListSeparator)
}
