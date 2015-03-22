package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// :(linux/unix) or ;(windows)
const _pathListSeparator = string(os.PathListSeparator)

type CompoundClassPathEntry struct {
	entries []ClassPathEntry
}

func newCompoundClassPathEntry(pathList string) *CompoundClassPathEntry {
	compoundEntry := &CompoundClassPathEntry{}

	for _, path := range strings.Split(pathList, _pathListSeparator) {
		if absPath, err := filepath.Abs(path); err == nil {
			entry := parseClassPathEntry(absPath)
			compoundEntry.addEntry(entry)
		} else {
			// todo
		}
	}

	return compoundEntry
}

func (self *CompoundClassPathEntry) readClassData(className string) (ClassPathEntry, []byte, error) {
	for _, entry := range self.entries {
		entry, data, err := entry.readClassData(className)
		if err == nil {
			return entry, data, nil
		}
	}

	// todo
	return nil, nil, classNotFoundErr
}

func (self *CompoundClassPathEntry) addEntry(entry ClassPathEntry) {
	_len := len(self.entries)
	if _len == cap(self.entries) {
		newEntries := make([]ClassPathEntry, _len, _len+8)
		copy(newEntries, self.entries)
		self.entries = newEntries
	}

	self.entries = append(self.entries, entry)
}

func (self *CompoundClassPathEntry) String() string {
	strs := make([]string, len(self.entries))

	for i, entry := range self.entries {
		strs[i] = entry.String()
	}

	return strings.Join(strs, _pathListSeparator)
}
