package classpath

type CompoundClassPathEntry struct {
	entries []ClassPathEntry
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
