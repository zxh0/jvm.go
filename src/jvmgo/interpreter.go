package main

import (
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
        inst := instructions.Decode(bcr)
        frame.SetNextPC(bcr.PC())

        inst.Execute(thread)
        thread.SetPC(frame.NextPC())
    }
}
