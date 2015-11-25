package interpreter

import (
	"fmt"

	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/jerrors"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// todo
func Loop(thread *rtda.Thread) {
	threadObj := thread.JThread()
	isDaemon := threadObj != nil && threadObj.GetFieldValue("daemon", "Z").(int32) == 1
	if !isDaemon {
		nonDaemonThreadStart()
	}

	_loop(thread)

	// terminate thread
	threadObj = thread.JThread()
	threadObj.Monitor().NotifyAll()
	if !isDaemon {
		nonDaemonThreadStop()
	}
}

/*func _loop(thread *rtda.Thread) {
	defer _catchErr(thread) // todo

	decoder := instructions.NewDecoder()
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		code := frame.Method().Code()
		inst, nextPC := decoder.Decode(code, pc)
		frame.SetNextPC(nextPC)

		// execute
		//_logInstruction(frame, inst)
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}*/

func _loop(thread *rtda.Thread) {
	defer _catchErr(thread) // todo

	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// fetch instruction
		method := frame.Method()
		if method.Instructions == nil {
			method.Instructions = decodeMethod(method.Code())
		}
		insts := method.Instructions.([]base.Instruction)
		inst := insts[pc]
		instCount := len(insts)

		// update nextPC
		for {
			pc++
			if pc >= instCount || insts[pc] != nil {
				break
			}
		}
		frame.SetNextPC(pc)

		// execute instruction
		//_logInstruction(frame, inst)
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
		fmt.Printf(">> line:%4d pc:%4d %v.%v%v \n",
			lineNum, frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func _logInstruction(frame *rtda.Frame, inst base.Instruction) {
	thread := frame.Thread()
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := thread.PC()

	if method.IsStatic() {
		fmt.Printf("[instruction] thread:%p %v.%v() #%v %T %v\n",
			thread, className, methodName, pc, inst, inst)
	} else {
		fmt.Printf("[instruction] thread:%p %v#%v() #%v %T %v\n",
			thread, className, methodName, pc, inst, inst)
	}
}
