package lang

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_throwable(fillInStackTrace, "fillInStackTrace", "(I)Ljava/lang/Throwable;")
	_throwable(getStackTraceElement, "getStackTraceElement", "(I)Ljava/lang/StackTraceElement;")
	_throwable(getStackTraceDepth, "getStackTraceDepth", "()I")
}

func _throwable(method native.Method, name, desc string) {
	native.Register("java/lang/Throwable", name, desc, method)
}

type StackTraceElement struct {
	declaringClass string
	methodName     string
	fileName       string
	lineNumber     int
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.GetRefVar(0)

	frame.PushRef(this)

	stes := createStackTraceElements(this, frame)
	this.Extra = stes
}

func createStackTraceElements(tObj *heap.Object, frame *rtda.Frame) []*StackTraceElement {
	thread := frame.Thread
	depth := thread.StackDepth()

	// skip unrelated frames
	i := uint(1)
	for k := tObj.Class; k != nil; k = k.SuperClass {
		i++
	}
	if thread.TopFrameN(i).Method.Name == "<athrow>" {
		i++
	}

	stes := make([]*StackTraceElement, 0, depth)
	for ; i < depth; i++ {
		frameN := thread.TopFrameN(i)
		methodN := frameN.Method
		classN := methodN.Class
		if classN.Name != "~shim" { // skip shim frame
			lineNumber := methodN.GetLineNumber(frameN.NextPC - 1)
			ste := &StackTraceElement{
				declaringClass: classN.NameJlsFormat(),
				methodName:     methodN.Name,
				fileName:       classN.SourceFile,
				lineNumber:     lineNumber,
			}
			stes = append(stes, ste)
		}
	}

	return stes
}

// native int getStackTraceDepth();
// ()I
func getStackTraceDepth(frame *rtda.Frame) {
	this := frame.GetRefVar(0)

	stes := this.Extra.([]*StackTraceElement)
	depth := int32(len(stes))

	frame.PushInt(depth)
}

// native StackTraceElement getStackTraceElement(int index);
// (I)Ljava/lang/StackTraceElement;
func getStackTraceElement(frame *rtda.Frame) {
	this := frame.GetRefVar(0)
	index := frame.GetIntVar(1)

	stes := this.Extra.([]*StackTraceElement)
	ste := stes[index]

	steObj := createStackTraceElementObj(ste, frame)
	frame.PushRef(steObj)
}

func createStackTraceElementObj(ste *StackTraceElement, frame *rtda.Frame) *heap.Object {
	declaringClass := frame.GetRuntime().JSFromGoStr(ste.declaringClass)
	methodName := frame.GetRuntime().JSFromGoStr(ste.methodName)
	fileName := frame.GetRuntime().JSFromGoStr(ste.fileName)
	lineNumber := int32(ste.lineNumber)

	/*
	   public StackTraceElement(String declaringClass, String methodName,
	           String fileName, int lineNumber)
	*/
	steClass := frame.GetClassLoader().LoadClass("java/lang/StackTraceElement")
	steObj := steClass.NewObj()
	// todo: call <init>
	steObj.SetFieldValue("declaringClass", "Ljava/lang/String;", heap.NewRefSlot(declaringClass))
	steObj.SetFieldValue("methodName", "Ljava/lang/String;", heap.NewRefSlot(methodName))
	steObj.SetFieldValue("fileName", "Ljava/lang/String;", heap.NewRefSlot(fileName))
	steObj.SetFieldValue("lineNumber", "I", heap.NewIntSlot(lineNumber))

	return steObj
}
