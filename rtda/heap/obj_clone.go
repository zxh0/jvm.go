package heap

func (self *Object) Clone() *Object {
	fields2 := self._cloneFields()
	var extra2 interface{} = nil // todo

	return newObj(self.class, fields2, extra2)
}

func (self *Object) _cloneFields() interface{} {
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
	case []*Object:
		fields := self.fields.([]*Object)
		fields2 := make([]*Object, len(fields))
		copy(fields2, fields)
		return fields2
	default: // []interface{}
		fields := self.fields.([]interface{})
		fields2 := make([]interface{}, len(fields))
		copy(fields2, fields)
		return fields2
	}
}
