package lang

import (
// "fmt"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _throwable(fillInStackTrace,        "fillInStackTrace",     "(I)Ljava/lang/Throwable;")
    _throwable(getStackTraceElement,    "getStackTraceElement", "(I)Ljava/lang/StackTraceElement;")
    _throwable(getStackTraceDepth,      "getStackTraceDepth",   "()I")
}

func _throwable(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Throwable", name, desc, method)
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopInt() // dummy
    this := stack.PopRef() // this
    stack.PushRef(this)

//     thread := frame.Thread()
//     depth := thread.StackDepth()
//     
//     steArrObj := steClass.NewArray(depth)
//     this.SetFieldValue("stackTrace", "[Ljava/lang/StackTraceElement;", steArrObj)
// st:= this.GetFieldValue("stackTrace", "[Ljava/lang/StackTraceElement;").(*rtc.Obj)
// fmt.Printf("#@@@@@:%v\n", st)
//     stes := steArrObj.Fields().([]*rtc.Obj)
//     for i := uint(0); i < depth; i++ {
//         frameN := thread.TopFrameN(i)
//         stes[i] = createStackTraceElement(frameN, steClass)
//     }
}

// native int getStackTraceDepth();
// ()I
func getStackTraceDepth(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopRef() // this

    depth := frame.Thread().StackDepth()
    stack.PushInt(int32(depth))
}

// native StackTraceElement getStackTraceElement(int index);
// (I)Ljava/lang/StackTraceElement;
func getStackTraceElement(frame *rtda.Frame) {
    stack := frame.OperandStack()
    index := uint(stack.PopInt())
    stack.PopRef() // this

    frameN := frame.Thread().TopFrameN(index)
    steObj := createStackTraceElement(frameN)
    stack.PushRef(steObj)
}

func createStackTraceElement(frame *rtda.Frame) (*rtc.Obj) {
    method := frame.Method()
    class := method.Class()

    declaringClass := rtda.NewJString(class.Name(), frame)
    methodName := rtda.NewJString(method.Name(), frame)
    fileName := rtda.NewJString(class.SourceFile(), frame)
    lineNumber := int32(-1) // todo

    /*
    public StackTraceElement(String declaringClass, String methodName,
            String fileName, int lineNumber)
    */
    steClass := frame.GetClassLoader().LoadClass("java/lang/StackTraceElement")
    ste := steClass.NewObj()
    ste.SetFieldValue("declaringClass", "Ljava/lang/String;",   declaringClass)
    ste.SetFieldValue("methodName",     "Ljava/lang/String;",   methodName)
    ste.SetFieldValue("fileName",       "Ljava/lang/String;",   fileName)
    ste.SetFieldValue("lineNumber",     "I",                    lineNumber)

    return ste
}
