package class

import (
	"fmt"
	"sync"
)

// object
type Obj struct {
	class   *Class
	fields  interface{} // []interface{} for Object, []int32 for int[] ...
	extra   interface{} // remember some important things from golnag
	monitor *Monitor
	lock    *sync.RWMutex // state lock
}

func newObj(class *Class, fields, extra interface{}) *Obj {
	return &Obj{class, fields, extra, newMonitor(), &sync.RWMutex{}}
}

func (self *Obj) String() string {
	return fmt.Sprintf("{Obj@%p class:%v extra:%v}",
		self, self.class, self.extra)
}

// getters & setters
func (self *Obj) Class() *Class {
	return self.class
}
func (self *Obj) Monitor() *Monitor {
	return self.monitor
}
func (self *Obj) Fields() interface{} {
	return self.fields
}
func (self *Obj) Extra() interface{} {
	return self.extra
}
func (self *Obj) SetExtra(extra interface{}) {
	self.extra = extra
}

func (self *Obj) GetPrimitiveDescriptor() string {
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
func (self *Obj) initFields() {
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
func (self *Obj) LockState() {
	self.lock.Lock()
}
func (self *Obj) UnlockState() {
	self.lock.Unlock()
}
func (self *Obj) RLockState() {
	self.lock.RLock()
}
func (self *Obj) RUnlockState() {
	self.lock.RUnlock()
}

// reflection
func (self *Obj) GetFieldValue(fieldName, fieldDescriptor string) interface{} {
	field := self.class.GetInstanceField(fieldName, fieldDescriptor)
	return field.GetValue(self)
}
func (self *Obj) SetFieldValue(fieldName, fieldDescriptor string, value interface{}) {
	field := self.class.GetInstanceField(fieldName, fieldDescriptor)
	field.PutValue(self, value)
}
