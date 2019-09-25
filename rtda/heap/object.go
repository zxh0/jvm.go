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

func (obj *Object) String() string {
	return fmt.Sprintf("{Object@%p class:%v extra:%v}",
		obj, obj.class, obj.extra)
}

// getters & setters
func (obj *Object) Class() *Class {
	return obj.class
}
func (obj *Object) Monitor() *Monitor {
	return obj.monitor
}
func (obj *Object) Fields() interface{} {
	return obj.fields
}
func (obj *Object) Extra() interface{} {
	return obj.extra
}
func (obj *Object) SetExtra(extra interface{}) {
	obj.extra = extra
}

func (obj *Object) GetPrimitiveDescriptor() string {
	switch obj.class.name {
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
func (obj *Object) initFields() {
	fields := obj.fields.([]interface{})
	for class := obj.class; class != nil; class = class.superClass {
		for _, f := range class.fields {
			if !f.IsStatic() {
				fields[f.slotId] = f.defaultValue()
			}
		}
	}
}

// state lock
func (obj *Object) LockState() {
	obj.lock.Lock()
}
func (obj *Object) UnlockState() {
	obj.lock.Unlock()
}
func (obj *Object) RLockState() {
	obj.lock.RLock()
}
func (obj *Object) RUnlockState() {
	obj.lock.RUnlock()
}

// reflection
func (obj *Object) GetFieldValue(fieldName, fieldDescriptor string) interface{} {
	field := obj.class.GetInstanceField(fieldName, fieldDescriptor)
	return field.GetValue(obj)
}
func (obj *Object) SetFieldValue(fieldName, fieldDescriptor string, value interface{}) {
	field := obj.class.GetInstanceField(fieldName, fieldDescriptor)
	field.PutValue(obj, value)
}
