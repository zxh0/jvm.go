package interpreter

import (
	"fmt"
	"jvmgo/jvm/instructions"
	"jvmgo/jvm/jerrors"
	"jvmgo/jvm/keepalive"
	"jvmgo/jvm/rtda"
)

// todo
func Loop(thread *rtda.Thread) {
	threadObj := thread.JThread()
	isDaemon := threadObj != nil && threadObj.GetFieldValue("daemon", "Z").(int32) == 1
	if !isDaemon {
		keepalive.NonDaemonThreadStart()
	}

	_loop(thread)

	// terminate thread
	threadObj = thread.JThread()
	threadObj.Monitor().NotifyAll()
	if !isDaemon {
		keepalive.NonDaemonThreadStop()
	}
}

func _loop(thread *rtda.Thread) {
	defer _catchErr(thread) // todo

	decoder := instructions.NewInstructionDecoder()
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		code := frame.Method().Code()
		_, inst, nextPC := decoder.Decode(code, pc)
		frame.SetNextPC(nextPC)

		// execute
		//_logInstruction(frame, opcode, inst)
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

// todo
func _catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		if err, ok := r.(jerrors.ClassNotFoundError); ok {
			thread.ThrowClassNotFoundException(err.Error())
			_loop(thread)
			return
		}

		_logFrames(thread)

		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
			panic(err.Error())
		} else {
			panic(err.Error())
		}
	}
}

func _logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		lineNum := method.GetLineNumber(frame.NextPC())
		fmt.Printf(">> %v.%v%v bytecode:%v line:%v \n", className, method.Name(), method.Descriptor(),
			frame.NextPC(), lineNum)
	}
}

func _logInstruction(frame *rtda.Frame, opcode uint8, inst instructions.Instruction) {
	thread := frame.Thread()
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := thread.PC()

	if method.IsStatic() {
		fmt.Printf("[instruction] thread:%p %v.%v() #%v 0x%x %v\n", thread, className, methodName, pc, opcode, inst)
	} else {
		fmt.Printf("[instruction] thread:%p %v#%v() #%v 0x%x %v\n", thread, className, methodName, pc, opcode, inst)
	}
}
