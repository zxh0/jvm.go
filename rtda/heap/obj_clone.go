package heap

import (
	"fmt"
)

func (obj *Object) Clone() *Object {
	fields2 := obj._cloneFields()
	var extra2 interface{} = nil // todo

	return newObj(obj.class, fields2, extra2)
}

func (obj *Object) _cloneFields() interface{} {
	switch fields := obj.fields.(type) {
	case []int8:
		fields2 := make([]int8, len(fields))
		copy(fields2, fields)
		return fields2
	case []int16:
		fields2 := make([]int16, len(fields))
		copy(fields2, fields)
		return fields2
	case []uint16:
		fields2 := make([]uint16, len(fields))
		copy(fields2, fields)
		return fields2
	case []int32:
		fields2 := make([]int32, len(fields))
		copy(fields2, fields)
		return fields2
	case []int64:
		fields2 := make([]int64, len(fields))
		copy(fields2, fields)
		return fields2
	case []float32:
		fields2 := make([]float32, len(fields))
		copy(fields2, fields)
		return fields2
	case []float64:
		fields2 := make([]float64, len(fields))
		copy(fields2, fields)
		return fields2
	case []*Object:
		fields2 := make([]*Object, len(fields))
		copy(fields2, fields)
		return fields2
	case []Slot:
		fields2 := make([]Slot, len(fields))
		copy(fields2, fields)
		return fields2
	default:
		panic(fmt.Errorf("not array: %v", obj))
	}
}
