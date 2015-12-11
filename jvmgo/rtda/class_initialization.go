package rtda

import (
	"unsafe"

	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

// see: jls8 12.4.2. Detailed Initialization Procedure
// http://docs.oracle.com/javase/specs/jls/se8/html/jls-12.html#jls-12.4.2
func initClass(thread *Thread, class *heap.Class) {
	// step 1
	initCond := class.InitCond()
	initCond.L.Lock()

	// step 2 & 3
	threadPtr := uintptr(unsafe.Pointer(thread))
	isInitializing, initThreadPtr := class.IsBeingInitialized()
	if isInitializing {
		if initThreadPtr != threadPtr {
			initCond.Wait()
		} else {
			initCond.L.Unlock()
			return
		}
	}

	// step 4
	if class.IsFullyInitialized() {
		initCond.L.Unlock()
		return
	}

	// step 5
	if class.IsInitializationFailed() {
		initCond.L.Unlock()
		panic("NoClassDefFoundError") // todo
	}

	// step 6
	class.MarkBeingInitialized(threadPtr)
	initCond.L.Unlock()
	initConstantStaticFields(class)

	// step 7
	defer initSuperClass(thread, class)

	// step 8
	// todo

	// step 9 & 10
	callClinit(thread, class)

	// step 11 & 12
	// todo
}

func initSuperClass(thread *Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && superClass.InitializationNotStarted() {
			initClass(thread, superClass)
		}
	}
}

func callClinit(thread *Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit == nil {
		clinit = heap.ReturnMethod() // just do nothing
	}

	// exec <clinit>
	newFrame := thread.NewFrame(clinit)
	newFrame.SetOnPopAction(func() {
		// step 10
		initSucceeded(class)
	})
	thread.PushFrame(newFrame)
}

// step 10
func initSucceeded(class *heap.Class) {
	initCond := class.InitCond()
	initCond.L.Lock()
	defer initCond.L.Unlock()

	class.MarkFullyInitialized()
	class.InitCond().Broadcast()
}

// todo
func initConstantStaticFields(class *heap.Class) {
	cp := class.ConstantPool()

	for _, field := range class.Fields() {
		if field.IsStatic() && field.IsFinal() {
			kValIndex := uint(field.ConstValueIndex())
			if kValIndex > 0 {
				slotId := field.SlotId()
				staticSlots := class.StaticFieldSlots()
				switch field.Descriptor() {
				case "Z", "B", "C", "S", "I":
					staticSlots[slotId] = cp.GetConstant(kValIndex).(int32)
				case "J":
					staticSlots[slotId] = cp.GetConstant(kValIndex).(int64)
				case "F":
					staticSlots[slotId] = cp.GetConstant(kValIndex).(float32)
				case "D":
					staticSlots[slotId] = cp.GetConstant(kValIndex).(float64)
				case "Ljava/lang/String;":
					staticSlots[slotId] = JString(cp.GetConstant(kValIndex).(string))
				}
			}
		}
	}
}
