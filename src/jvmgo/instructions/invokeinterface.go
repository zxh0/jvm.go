package instructions

import (
    //"log"
    //. "jvmgo/any"
    //"jvmgo/native"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Invoke interface method
type invokeinterface struct {
    index   uint16
    count   uint8 // unused
    // 0
}

func (self *invokeinterface) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
    self.count = bcr.readUint8()
    bcr.readUint8() // must be 0
}

func (self *invokeinterface) Execute(frame *rtda.Frame) {
    thread := frame.Thread()
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    c := cp.GetConstant(uint(self.index))
    cMethodRef := c.(*rtc.ConstantInterfaceMethodref)
    ref := stack.Top(cMethodRef.ArgCount())
    if ref == nil {
        panic("NPE")
    }

    method := cMethodRef.VirtualMethod(ref.(*rtc.Obj))
    if method.IsNative() {
        nativeMethod := method.NativeMethod().(func(*rtda.Frame))
        nativeMethod(frame)
    } else {
        thread.InvokeMethod(method)
    }
}
