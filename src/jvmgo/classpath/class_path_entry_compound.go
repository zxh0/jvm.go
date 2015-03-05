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
