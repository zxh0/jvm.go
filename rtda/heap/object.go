package heap

import (
	"fmt"
	"reflect"
	"sync"
)

// object
type Object struct {
	Class   *Class
	Fields  interface{} // []Slot for Object, []int32 for int[] ...
	Extra   interface{} // remember some important things from Golang
	Monitor *Monitor
	lock    *sync.RWMutex // state lock
}

func newObj(class *Class, fields, extra interface{}) *Object {
	return &Object{class, fields, extra, newMonitor(), &sync.RWMutex{}}
}

func (obj *Object) String() string {
	return fmt.Sprintf("{Object@%p class:%v extra:%v}",
		obj, obj.Class, obj.Extra)
}

// todo
func (obj *Object) initFields() {
	fields := obj.Fields.([]Slot)
	for class := obj.Class; class != nil; class = class.SuperClass {
		for _, f := range class.Fields {
			if !f.IsStatic() {
				fields[f.SlotId] = EmptySlot // TODO
			}
		}
	}
}

// state lock
func (obj *Object) LockState()    { obj.lock.Lock() }
func (obj *Object) UnlockState()  { obj.lock.Unlock() }
func (obj *Object) RLockState()   { obj.lock.RLock() }
func (obj *Object) RUnlockState() { obj.lock.RUnlock() }

// reflection
func (obj *Object) GetFieldValue(fieldName, fieldDescriptor string) Slot {
	field := obj.Class.GetInstanceField(fieldName, fieldDescriptor)
	return field.GetValue(obj)
}
func (obj *Object) SetFieldValue(fieldName, fieldDescriptor string, value Slot) {
	field := obj.Class.GetInstanceField(fieldName, fieldDescriptor)
	field.PutValue(obj, value)
}

func (obj *Object) Clone() *Object {
	fields1 := reflect.ValueOf(obj.Fields)
	fields2 := reflect.MakeSlice(fields1.Type(), fields1.Len(), fields1.Len())
	reflect.Copy(fields2, fields1)
	var extra2 interface{} = nil // todo
	return newObj(obj.Class, fields2.Interface(), extra2)
}

func (obj *Object) GetGoClass() *Class {
	return obj.Extra.(*Class)
}
