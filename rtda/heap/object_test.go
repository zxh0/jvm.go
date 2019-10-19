package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClone(t *testing.T) {
	obj1 := newObj(nil, []int8{1, 2, 3}, nil)
	obj2 := obj1.Clone()
	obj2.Fields.([]int8)[1] = 9
	require.Equal(t, []int8{1, 2, 3}, obj1.Fields)
	require.Equal(t, []int8{1, 9, 3}, obj2.Fields)
}

func TestArrayLength(t *testing.T) {
	arr := newObj(nil, []int8{1, 2, 3, 4, 5}, nil)
	require.Equal(t, int32(5), arr.ArrayLength())
}

func TestArrayCopy(t *testing.T) {
	arr1 := newObj(nil, []int8{1, 2, 3, 4, 5}, nil)
	arr2 := newObj(nil, []int8{6, 7, 8, 9, 0}, nil)
	ArrayCopy(arr2, arr1, 1, 2, 3)
	require.Equal(t, []int8{6, 7, 8, 9, 0}, arr2.Fields)
	require.Equal(t, []int8{1, 2, 7, 8, 9}, arr1.Fields)
}
