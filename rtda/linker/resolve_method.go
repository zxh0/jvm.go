package linker

import (
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vm"
)

func ResolveMethod(loader *heap.ClassLoader, ref *heap.ConstantMethodRef) *heap.Method {
	if ref.ResolvedMethod == nil {
		ref.ResolvedClass = loader.LoadClass(ref.ClassName)
		if !ref.IsInterface {
			ref.ResolvedMethod = lookupMethod(ref.ResolvedClass, ref.Name, ref.Descriptor)
		} else {
			ref.ResolvedMethod = lookupInterfaceMethod(
				loader.JLObjectClass(), ref.ResolvedClass, ref.Name, ref.Descriptor)
		}
		if ref.ResolvedMethod == nil {
			panic(vm.NewNoSuchMethodError(ref.String()))
		}
		// TODO: apply access control
	}
	return ref.ResolvedMethod
}

// https://docs.oracle.com/javase/specs/jvms/se13/html/jvms-5.html#jvms-5.4.3.3
func lookupMethod(class *heap.Class, name, descriptor string) *heap.Method {
	// 1. If C is an interface, method resolution throws an IncompatibleClassChangeError.
	if class.IsInterface() {
		panic("IncompatibleClassChangeError") // TODO
	}

	// 2. Otherwise, method resolution attempts to locate the referenced method in C
	// and its superclasses:
	if method := lookupMethodInClasses(class, name, descriptor); method != nil {
		return method
	}

	// 3. Otherwise, method resolution attempts to locate the referenced method in the
	// superinterfaces of the specified class C:
	return lookupMethodInInterfaces(class, name, descriptor)
}

func lookupMethodInClasses(class *heap.Class, name, descriptor string) *heap.Method {
	// TODO: Signature Polymorphic Methods

	// Otherwise,if C declares a method with the name and descriptor
	// specified by the method reference, method lookup succeeds.
	if method := class.GetDeclaredMethod(name, descriptor); method != nil {
		return method
	}

	// Otherwise, if C has a superclass, step 2 of method resolution
	// is recursively invoked on the direct superclass of C.
	if class.SuperClass != nil {
		return lookupMethodInClasses(class.SuperClass, name, descriptor)
	}

	return nil
}

// TODO
func lookupMethodInInterfaces(class *heap.Class, name, descriptor string) *heap.Method {
	for _, iface := range class.Interfaces {
		for _, method := range iface.Methods {
			if method.Name == name && method.Descriptor == descriptor &&
				!method.IsPrivate() && !method.IsStatic() {

				return method
			}
		}
		if method := lookupMethodInInterfaces(iface, name, descriptor); method != nil {
			return method
		}
	}
	return nil
}

// https://docs.oracle.com/javase/specs/jvms/se13/html/jvms-5.html#jvms-5.4.3.4
func lookupInterfaceMethod(objClass, class *heap.Class, name, descriptor string) *heap.Method {
	// 1. If C is not an interface, interface method resolution throws an IncompatibleClassChangeError.
	if !class.IsInterface() {
		panic("IncompatibleClassChangeError") // TODO
	}

	// 2. Otherwise, if C declares a method with the name and descriptor specified by
	// the interface method reference, method lookup succeeds.
	if method := class.GetDeclaredMethod(name, descriptor); method != nil {
		return method
	}

	// 3. Otherwise, if the class Object declares a method with the name and descriptor
	// specified by the interface method reference, which has its ACC_PUBLIC flag set
	// and does not have its ACC_STATIC flag set, method lookup succeeds.
	if method := objClass.GetDeclaredMethod(name, descriptor); method != nil {
		if method.IsPublic() && !method.IsStatic() {
			return method
		}
	}

	// 4. & 5.
	return lookupMethodInInterfaces(class, name, descriptor)
}

/*
A maximally-specific superinterface method of a class or interface C for a particular
method name and descriptor is any method for which all of the following are true:
  1. The method is declared in a superinterface (direct or indirect) of C.
  2. The method is declared with the specified name and descriptor.
  3. The method has neither its ACC_PRIVATE flag nor its ACC_STATIC flag set.
  4. Where the method is declared in interface I, there exists no other maximally-specific
     superinterface method of C with the specified name and descriptor that is declared in
     a subinterface of I.
*/
func lookupMSSIMethod(class *heap.Class, methodName, descriptor string) *heap.Method {
	// TODO
	return nil
}
