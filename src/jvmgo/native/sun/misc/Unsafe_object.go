package misc

import (
    //"unsafe"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _unsafe(getObject,          "getObject",            "(Ljava/lang/reflect/Field;)J")
    _unsafe(getObjectVolatile,  "getObjectVolatile",    "(Ljava/lang/Object;J)Ljava/lang/Object;")
}

// public native Object getObject(Object o, long offset);
// (Ljava/lang/Object;J)Ljava/lang/Object;
func getObject(frame *rtda.Frame) {
    stack := frame.OperandStack()
    offset := stack.PopLong()
    obj := stack.PopRef()
    stack.PopRef() // this

    fields := obj.Fields().([]Any)
    fieldVal := fields[offset].(*rtc.Obj)
    stack.PushRef(fieldVal)
}

// public native Object getObjectVolatile(Object o, long offset);
//(Ljava/lang/Object;J)Ljava/lang/Object;
func getObjectVolatile(frame *rtda.Frame) {
    getObject(frame) // todo    
}
