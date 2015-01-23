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
        opcode, inst := instructions.Decode(bcr)
        frame.SetNextPC(bcr.PC())

        // execute
        log.Printf("exec instruction: 0x%x %v", opcode, inst)
        inst.Execute(thread)
        if !thread.IsStackEmpty() {
            topFrame := thread.TopFrame()
            thread.SetPC(topFrame.NextPC())
        } else {
            break;
        }
    }
}
