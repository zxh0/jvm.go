package class

import . "jvmgo/any"

func GetFieldValue(obj *Obj, fieldName, fieldDescriptor string) Any {
    field := obj.class.GetField(fieldName, fieldDescriptor)
    return field.GetValue(obj)
}
