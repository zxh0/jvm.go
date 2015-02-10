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
func getObject(frame *rtda.Frame, x int) {
    vars := frame.LocalVars()
    // vars.GetRef(0) // this
    obj := vars.GetRef(1)
    offset := vars.GetLong(2)

    fields := obj.Fields().([]Any)
    fieldVal := fields[offset].(*rtc.Obj)

    stack := frame.OperandStack()
    stack.PushRef(fieldVal)
}

// public native Object getObjectVolatile(Object o, long offset);
//(Ljava/lang/Object;J)Ljava/lang/Object;
func getObjectVolatile(frame *rtda.Frame, x int) {
    getObject(frame, x) // todo    
}
