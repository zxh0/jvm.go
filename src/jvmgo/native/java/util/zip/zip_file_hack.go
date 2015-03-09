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
		return key, err
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
