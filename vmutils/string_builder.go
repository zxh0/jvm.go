package vmutils

import "strings"

type StringBuilder struct {
	ss []string
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{ss: make([]string, 0, 8)}
}

func (sb *StringBuilder) Append(ss ...string) {
	sb.ss = append(sb.ss, ss...)
}

func (sb *StringBuilder) String() string {
	return strings.Join(sb.ss, "")
}
