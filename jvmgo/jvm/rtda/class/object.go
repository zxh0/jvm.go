package class

import (
	"fmt"
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"sync"
)

// object
type Obj struct {
	class   *Class
	fields  Any // []Any for Object, []int32 for int[] ...
	extra   Any // todo
	monitor *Monitor
	lock    *sync.RWMutex // state lock
}

func newObj(class *Class, fields, extra Any) *Obj {
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
func (self *Obj) Fields() Any {
	return self.fields
}
func (self *Obj) Extra() Any {
	return self.extra
}
func (self *Obj) SetExtra(extra Any) {
	self.extra = extra
}

// todo
func (self *Obj) initFields() {
	fields := self.fields.([]Any)
	for class := self.class; class != nil; class = class.superClass {
		for _, f := range class.fields {
			if !f.IsStatic() {
				fields[f.slot] = f.defaultValue()
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
func (self *Obj) GetFieldValue(fieldName, fieldDescriptor string) Any {
	field := self.class.GetInstanceField(fieldName, fieldDescriptor)
	return field.GetValue(self)
}
func (self *Obj) SetFieldValue(fieldName, fieldDescriptor string, value Any) {
	field := self.class.GetInstanceField(fieldName, fieldDescriptor)
	field.PutValue(self, value)
}
