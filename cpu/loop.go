package cpu

import (
	"fmt"

	"github.com/zxh0/jvm.go/instructions"
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vm"
)

func ExecMethod(thread *rtda.Thread, method *heap.Method, args []heap.Slot) heap.Slot {
	shimFrame := rtda.NewShimFrame(thread, args)
	thread.PushFrame(shimFrame)
	thread.InvokeMethod(method)

	debug := thread.VMOptions.XDebugInstr
	defer _catchErr(thread) // todo

	for {
		frame := thread.CurrentFrame()
		if frame == shimFrame {
			thread.PopFrame()
			if frame.IsStackEmpty() {
				return heap.EmptySlot
			} else {
				return frame.Pop()
			}
		}

		pc := frame.NextPC
		thread.PC = pc

		// fetch instruction
		instr, nextPC := fetchInstruction(frame.Method, pc)
		frame.NextPC = nextPC

		// execute instruction
		instr.Execute(frame)
		if debug {
			_logInstruction(frame, instr)
		}
	}
}

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
	debug := thread.VMOptions.XDebugInstr
	defer _catchErr(thread) // todo

	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC
		thread.PC = pc

		// fetch instruction
		instr, nextPC := fetchInstruction(frame.Method, pc)
		frame.NextPC = nextPC

		// execute instruction
		instr.Execute(frame)
		if debug {
			_logInstruction(frame, instr)
		}
		if thread.IsStackEmpty() {
			break
		}
	}
}

func fetchInstruction(method *heap.Method, pc int) (base.Instruction, int) {
	if method.Instructions == nil {
		method.Instructions = instructions.Decode(method.Code)
	}

	instrs := method.Instructions.([]base.Instruction)
	instr := instrs[pc]

	// calc nextPC
	pc++
	for pc < len(instrs) && instrs[pc] == nil {
		pc++
	}

	return instr, pc
}

// todo
func _catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		if err, ok := r.(vm.ClassNotFoundError); ok {
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
