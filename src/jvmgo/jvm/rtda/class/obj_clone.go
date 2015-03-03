package class

import (
	. "jvmgo/any"
)

func (self *Obj) Clone() *Obj {
	fields2 := self._cloneFields()
	var extra2 Any = nil // todo

	return newObj(self.class, fields2, extra2)
}

func (self *Obj) _cloneFields() Any {
	switch self.fields.(type) {
	case []int8:
		fields := self.fields.([]int8)
		fields2 := make([]int8, len(fields))
		copy(fields2, fields)
		return fields2
	case []int16:
		fields := self.fields.([]int16)
		fields2 := make([]int16, len(fields))
		copy(fields2, fields)
		return fields2
	case []uint16:
		fields := self.fields.([]uint16)
		fields2 := make([]uint16, len(fields))
		copy(fields2, fields)
		return fields2
	case []int32:
		fields := self.fields.([]int32)
		fields2 := make([]int32, len(fields))
		copy(fields2, fields)
		return fields2
	case []int64:
		fields := self.fields.([]int64)
		fields2 := make([]int64, len(fields))
		copy(fields2, fields)
		return fields2
	case []float32:
		fields := self.fields.([]float32)
		fields2 := make([]float32, len(fields))
		copy(fields2, fields)
		return fields2
	case []float64:
		fields := self.fields.([]float64)
		fields2 := make([]float64, len(fields))
		copy(fields2, fields)
		return fields2
	case []*Obj:
		fields := self.fields.([]*Obj)
		fields2 := make([]*Obj, len(fields))
		copy(fields2, fields)
		return fields2
	default: // []Any
		fields := self.fields.([]Any)
		fields2 := make([]Any, len(fields))
		copy(fields2, fields)
		return fields2
	}
}
