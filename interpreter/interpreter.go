package interpreter

import (
	"fmt"

	"github.com/zxh0/jvm.go/instructions"
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vmerrors"
)

// todo
func Loop(thread *rtda.Thread) {
	threadObj := thread.JThread()
	isDaemon := threadObj != nil && threadObj.GetFieldValue("daemon", "Z").IntValue() == 1
	if !isDaemon {
		nonDaemonThreadStart()
	}

	_loop(thread)

	// terminate thread
	threadObj = thread.JThread()
	threadObj.Monitor.NotifyAll()
	if !isDaemon {
		nonDaemonThreadStop()
	}
}

func _loop(thread *rtda.Thread) {
	verbose := thread.VMOptions.VerboseInstr
	defer _catchErr(thread) // todo

	for {
		frame := thread.CurrentFrame()
		thread.PC = frame.NextPC
		frame.NextPC++

		// fetch & execute instruction
		instr := getInstruction(frame.Method, thread.PC)
		instr.Execute(frame)
		if verbose {
			_logInstruction(frame, instr)
		}
		if thread.IsStackEmpty() {
			break
		}
	}
}

func getInstruction(method *heap.Method, pc int) base.Instruction {
	if method.Instructions == nil {
		method.Instructions = instructions.Decode(method.Code, true)
	}
	return method.Instructions.([]base.Instruction)[pc]
}

// todo
func _catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		if err, ok := r.(vmerrors.ClassNotFoundError); ok {
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
		method := frame.Method
		className := method.Class.Name
		lineNum := method.GetLineNumber(frame.NextPC)
		fmt.Printf(">> line:%4d pc:%4d %v.%v%v \n",
			lineNum, frame.NextPC, className, method.Name, method.Descriptor)
	}
}

func _logInstruction(frame *rtda.Frame, instr base.Instruction) {
	thread := frame.Thread
	method := frame.Method
	className := method.Class.Name
	pc := thread.PC

	if method.IsStatic() {
		fmt.Printf("[instruction] thread:%p %v.%v() #%v %T %v\n",
			thread, className, method.Name, pc, instr, instr)
	} else {
		fmt.Printf("[instruction] thread:%p %v#%v() #%v %T %v\n",
			thread, className, method.Name, pc, instr, instr)
	}
}
