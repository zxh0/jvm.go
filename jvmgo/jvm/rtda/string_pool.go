package rtda

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

// home for interned Strings
var _stringPool = []StringItem{}

type StringItem struct {
	chars []uint16
	str   *rtc.Obj
}

func getInternedString(chars []uint16) *rtc.Obj {
	index := _binarySearch(chars)
	if index >= 0 {
		return _stringPool[index].str
	} else {
		return nil
	}
}

func InternString(chars []uint16, str *rtc.Obj) *rtc.Obj {
	index := _binarySearch(chars)
	if index >= 0 {
		return _stringPool[index].str
	} else {
		_insert(-index-1, chars, str)
		return str
	}
}

func _binarySearch(chars []uint16) int {
	low := 0
	high := len(_stringPool) - 1

	for low <= high {
		mid := (low + high) / 2
		midVal := _stringPool[mid].chars

		c := _compare(midVal, chars)
		if c < 0 {
			low = mid + 1
		} else if c > 0 {
			high = mid - 1
		} else {
			return mid // key found
		}
	}
	return -(low + 1) // key not found.
}

func _compare(a, b []uint16) int {
	aLen, bLen := len(a), len(b)
	minLen := _min(aLen, bLen)
	for i := 0; i < minLen; i++ {
		x, y := a[i], b[i]
		if x < y {
			return -1
		} else if x > y {
			return 1
		}
	}
	return aLen - bLen
}

func _min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func _insert(index int, chars []uint16, str *rtc.Obj) {
	poolLen := len(_stringPool)
	if poolLen == cap(_stringPool) {
		_expandPool()
	}

	_stringPool = _stringPool[:poolLen+1]
	src := _stringPool[index:poolLen]
	dst := _stringPool[index+1 : poolLen+1]
	copy(dst, src)

	_stringPool[index] = StringItem{chars, str}
}

func _expandPool() {
	poolLen := len(_stringPool)
	newPool := make([]StringItem, poolLen, poolLen+100)
	copy(newPool, _stringPool) // func copy(dst, src []T) int
	_stringPool = newPool
}
