package jutil

import (
	"strings"
)

func ReplaceAll(s, old, _new string) string {
	return strings.Replace(s, old, _new, -1)
}
