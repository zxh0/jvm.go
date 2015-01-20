package main

import (
    "jvmgo/rtda"
    "jvmgo/instructions"
)

func interpret() {

}

// todo
func loop(thread *rtda.Thread) {
    bcr := instructions.NewBytecodeReader()
    for !thread.IsStackEmpty() {
        bcr.SetPC(thread.PC())
        bcr.SetCode(thread.CurrentFrame().Method().Code())
        inst := instructions.Decode(bcr)
        inst.execute(thread)
    }
}

// func executeOneInstruction(frame *rtda.Frame) {
//     // todo
    
// }
