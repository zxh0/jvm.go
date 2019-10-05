package jimage

import (
	"strings"
)

func IsTreeInfoResource(path string) bool {
	return strings.HasPrefix(path, "/packages") ||
		strings.HasPrefix(path, "/modules")
}
