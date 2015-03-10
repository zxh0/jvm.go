package zip

import (
	gozip "archive/zip"
	"bytes"
)

var _map = map[int64]*gozip.ReadCloser{}

func openZip(name string) (int64, error) {
	rc, err := gozip.OpenReader(name)
	if err == nil {
		jzfile := _nextJzfile()
		_map[jzfile] = rc
		return jzfile, nil
	}
	return 0, err
}

func _nextJzfile() int64 {
	maxKey := int64(0)
	for key, _ := range _map {
		if key >= maxKey {
			maxKey = key
		}
	}
	return maxKey + 1
}

func getEntryCount(jzfile int64) int32 {
	if rc, ok := _map[jzfile]; ok {
		return int32(len(rc.File))
	}
	// todo
	return 0
}

func getJzentry(jzfile int64, entryIndex int32) int64 {
	return jzfile | (int64(entryIndex) << 32)
}

func getJzentry2(jzfile int64, name []byte) int64 {
	entryIndex := _getEntryIndex(jzfile, name)
	if entryIndex >= 0 {
		return getJzentry(jzfile, entryIndex)
	}
	return 0
}
func _getEntryIndex(jzfile int64, name []byte) int32 {
	if rc, ok := _map[jzfile]; ok {
		for i, f := range rc.File {
			if bytes.Equal(name, []byte(f.Name)) {
				return int32(i)
			}
		}
	}
	return -1
}

func getEntryFile(jzentry int64) *gozip.File {
	jzfile := jzentry & 0x0000FFFF
	entryIndex := jzentry >> 32

	if rc, ok := _map[jzfile]; ok {
		return rc.File[entryIndex]
	}
	// todo
	return nil
}

func readEntry(jzentry, pos int64, buf []byte) int {
	f := getEntryFile(jzentry)
	rc, err := f.Open()
	if err == nil {
		defer rc.Close()

		// todo
		if pos > 0 {
			s := make([]byte, pos)
			rc.Read(s)
		}

		n, err := rc.Read(buf)
		if err == nil || n > 0 {
			return n
		}
	}

	// todo
	return -1
}
