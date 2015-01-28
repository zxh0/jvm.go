package main

import (
    "log"
    "jvmgo/rtda"
    "jvmgo/instructions"
)

// todo
func loop(thread *rtda.Thread) {
    bcr := instructions.NewBytecodeReader()
    for {
        frame := thread.CurrentFrame() 

        // decode
        bcr.SetPC(thread.PC())
        bcr.SetCode(frame.Method().Code())
        _, inst := instructions.Decode(bcr)
        frame.SetNextPC(bcr.PC())

        // execute
        //pc := thread.PC()
        //logInstruction(frame, pc, opcode, inst)
        inst.Execute(thread)
        if !thread.IsStackEmpty() {
            topFrame := thread.TopFrame()
            thread.SetPC(topFrame.NextPC())
        } else {
            break;
        }
    }
}

func logInstruction(frame *rtda.Frame, pc int, opcode uint8, inst instructions.Instruction) {
    className := frame.Method().Class().Name()
    methodName := frame.Method().Name()
    log.Printf("exec: %v.%v #%v 0x%x %v", className, methodName, pc, opcode, inst)
}
