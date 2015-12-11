package heap

import (
	"fmt"
	"sync"
)

// object
type Object struct {
	class   *Class
	fields  interface{} // []interface{} for Object, []int32 for int[] ...
	extra   interface{} // remember some important things from golnag
	monitor *Monitor
	lock    *sync.RWMutex // state lock
}

func newObj(class *Class, fields, extra interface{}) *Object {
	return &Object{class, fields, extra, newMonitor(), &sync.RWMutex{}}
}

func (self *Object) String() string {
	return fmt.Sprintf("{Object@%p class:%v extra:%v}",
		self, self.class, self.extra)
}

// getters & setters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Monitor() *Monitor {
	return self.monitor
}
func (self *Object) Fields() interface{} {
	return self.fields
}
func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

func (self *Object) GetPrimitiveDescriptor() string {
	switch self.class.name {
	case "java/lang/Boolean":
		return "Z"
	case "java/lang/Byte":
		return "B"
	case "java/lang/Character":
		return "C"
	case "java/lang/Short":
		return "S"
	case "java/lang/Integer":
		return "I"
	case "java/lang/Long":
		return "J"
	case "java/lang/Float":
		return "F"
	case "java/lang/Double":
		return "D"
	default:
		return ""
	}
}

// todo
func (self *Object) initFields() {
	fields := self.fields.([]interface{})
	for class := self.class; class != nil; class = class.superClass {
		for _, f := range class.fields {
			if !f.IsStatic() {
				fields[f.slotId] = f.defaultValue()
			}
		}
	}
}

// state lock
func (self *Object) LockState() {
	self.lock.Lock()
}
func (self *Object) UnlockState() {
	self.lock.Unlock()
}
func (self *Object) RLockState() {
	self.lock.RLock()
}
func (self *Object) RUnlockState() {
	self.lock.RUnlock()
}

// reflection
func (self *Object) GetFieldValue(fieldName, fieldDescriptor string) interface{} {
	field := self.class.GetInstanceField(fieldName, fieldDescriptor)
	return field.GetValue(self)
}
func (self *Object) SetFieldValue(fieldName, fieldDescriptor string, value interface{}) {
	field := self.class.GetInstanceField(fieldName, fieldDescriptor)
	field.PutValue(self, value)
}
