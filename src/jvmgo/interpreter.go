package main

import (
    "log"
    "jvmgo/rtda"
    "jvmgo/instructions"
)

// todo
func loop(thread *rtda.Thread) {
    bcr := instructions.NewBytecodeReader()
    for !thread.IsStackEmpty() {
        frame := thread.CurrentFrame() 

        bcr.SetPC(thread.PC())
        bcr.SetCode(frame.Method().Code())
        opcode, inst := instructions.Decode(bcr)
        log.Printf("exec instruction: 0x%x %v", opcode, inst)
        frame.SetNextPC(bcr.PC())

        inst.Execute(thread)
        thread.SetPC(frame.NextPC())
    }
}
