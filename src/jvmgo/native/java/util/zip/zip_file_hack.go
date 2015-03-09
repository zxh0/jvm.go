package zip

import (
	gozip "archive/zip"
)

var _map = map[int64]*gozip.ReadCloser{}

func openZip(name string) (int64, error) {
	rc, err := gozip.OpenReader(name)
	if err == nil {
		key := _nextKey()
		_map[key] = rc
		return key, nil
	}
	return 0, err
}

func _nextKey() int64 {
	maxKey := int64(0)
	for key, _ := range _map {
		if key >= maxKey {
			maxKey = key
		}
	}
	return maxKey + 1
}

func getEntryCount(key int64) int32 {
	rc, ok := _map[key]
	if ok {
		return int32(len(rc.File))
	}
	// todo
	return 0
}
