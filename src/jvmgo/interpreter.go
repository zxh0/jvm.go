package main

import (
    "jvmgo/rtda"
    "jvmgo/instructions"
)

// todo
func loop(thread *rtda.Thread) {
    bcr := instructions.NewBytecodeReader()
    for !thread.IsStackEmpty() {
        pc := thread.PC()
        frame := thread.CurrentFrame() 
        code := frame.Method().Code()

        bcr.SetPC(pc)
        bcr.SetCode(code)
        inst := instructions.Decode(bcr)
        frame.SetNextPC(bcr.PC())

        inst.Execute(thread)
        thread.SetPC(frame.NextPC())
    }
}
