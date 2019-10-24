package module

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModuleInfo(t *testing.T) {
	bytes, err := ioutil.ReadFile("../test/testdata/java13/module-info.class")
	require.NoError(t, err)

	info := ParseModuleInfo(bytes)
	require.Equal(t, info, &Info{
		Name:    "hello.modules",
		Flags:   0,
		Version: "0.1",
		Requires: []Require{
			{
				Name:    "java.base",
				Flags:   0x8000,
				Version: "13.0.1",
			},
		},
		Exports: []Export{
			{
				Package: "hello",
				Flags:   0,
				To:      []string{},
			},
		},
		Opens:    []Open{},
		Uses:     []string{},
		Provides: []Provide{},
	})
}
