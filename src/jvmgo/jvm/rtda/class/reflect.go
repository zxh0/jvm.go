package class

import . "jvmgo/any"

func GetFieldValue(obj *Obj, fieldName, fieldDescriptor string) Any {
    class := obj.Class()
    field := class.GetField(fieldName, fieldDescriptor)
    value := field.GetValue(obj)
    return value
}
