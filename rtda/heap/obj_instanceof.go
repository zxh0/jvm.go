package heap

func (self *Object) IsInstanceOf(class *Class) bool {
	s, t := self.class, class
	return _checkcast(s, t)
}

// jvms8-6.5.checkcast
// jvms8-6.5.instanceof
func _checkcast(s, t *Class) bool {
	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			if !t.IsInterface() {
				return s.isSubClassOf(t)
			} else {
				return s.isImplements(t)
			}
		} else {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isSuperInterfaceOf(s)
			}
		}
	} else { // s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else { // t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || _checkcast(sc, tc)
		}
	}

	return false
}
