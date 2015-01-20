package main

import (
    "jvmgo/rtda"
    "jvmgo/instructions"
)

// todo
func loop(thread *rtda.Thread) {
    bcr := instructions.NewBytecodeReader()
    for !thread.IsStackEmpty() {
        code := thread.CurrentFrame().Method().Code()
        pc := thread.PC()

        bcr.SetPC(pc)
        bcr.SetCode(code)
        inst := instructions.Decode(bcr)

        inst.Execute(thread)
        // todo: correct?
        if pc == thread.PC() {
            // no branch
            pc = bcr.PC()
            thread.SetPC(pc)
        }
    }
}
