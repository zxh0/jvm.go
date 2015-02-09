package jvm

import (
    "log"
    "jvmgo/jvm/rtda"
    "jvmgo/jvm/instructions"
)

// todo
func loop(thread *rtda.Thread) {
    decoder := instructions.NewInstructionDecoder()
    for {
        frame := thread.CurrentFrame() 

        // decode
        decoder.SetPC(thread.PC())
        decoder.SetCode(frame.Method().Code())
        _, inst := decoder.Decode()
        frame.SetNextPC(decoder.PC())

        // execute
        //_logInstruction(frame, thread.PC(), opcode, inst)
        inst.Execute(frame)
        if !thread.IsStackEmpty() {
            topFrame := thread.TopFrame()
            thread.SetPC(topFrame.NextPC())
        } else {
            break;
        }
    }
}

func _logInstruction(frame *rtda.Frame, pc int, opcode uint8, inst instructions.Instruction) {
    method := frame.Method()
    methodName := method.Name()
    className := method.Class().Name()
    if method.IsStatic() {
        log.Printf("exec: %v.%v() #%v 0x%x %v", className, methodName, pc, opcode, inst)
    } else {
        log.Printf("exec: %v#%v() #%v 0x%x %v", className, methodName, pc, opcode, inst)
    }
}
