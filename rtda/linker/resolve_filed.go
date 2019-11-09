package linker

import (
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vm"
)

// https://docs.oracle.com/javase/specs/jvms/se13/html/jvms-5.html#jvms-5.4.3.2
func ResolveField(loader *heap.ClassLoader,
	ref *heap.ConstantFieldRef, static bool) *heap.Field {

	if ref.ResolvedField == nil {
		ref.ResolvedClass = loader.LoadClass(ref.ClassName)
		ref.ResolvedField = lookupField(ref.ResolvedClass, ref.Name, ref.Descriptor)
	}
	if ref.ResolvedField == nil {
		panic(vm.NewNoSuchFieldError(ref.String()))
	}
	if ref.ResolvedField.IsStatic() != static {
		panic(vm.NewIncompatibleClassChangeError(ref.ResolvedField.String()))
	}
	// TODO: apply access control
	return ref.ResolvedField
}

func lookupField(class *heap.Class, name, descriptor string) *heap.Field {
	// 1. If C declares a field with the name and descriptor specified by the field reference,
	// field lookup succeeds. The declared field is the result of the field lookup.
	if field := class.GetDeclaredField(name, descriptor); field != nil {
		return field
	}

	// 2. Otherwise, field lookup is applied recursively to the direct superinterfaces
	// of the specified class or interface C.
	if field := lookupFieldInInterfaces(class, name, descriptor); field != nil {
		return field
	}

	// 3. Otherwise, if C has a superclass S, field lookup is applied recursively to S.
	if class.SuperClass != nil {
		return lookupField(class.SuperClass, name, descriptor)
	}

	// 4. Otherwise, field lookup fails.
	return nil
}

func lookupFieldInInterfaces(class *heap.Class, name, descriptor string) *heap.Field {
	for _, iface := range class.Interfaces {
		for _, field := range iface.Fields {
			if field.Name == name && field.Descriptor == descriptor {
				return field
			}
		}
		if field := lookupFieldInInterfaces(iface, name, descriptor); field != nil {
			return field
		}
	}
	return nil
}
