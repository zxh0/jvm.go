package vmutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFuncName(t *testing.T) {
	require.Equal(t, "GetFuncName", GetFuncName(GetFuncName))
	require.Equal(t, "TestGetFuncName", GetFuncName(TestGetFuncName))
}
