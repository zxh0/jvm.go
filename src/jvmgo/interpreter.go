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
        currentFrame := thread.CurrentFrame()
        bcr.SetPC(thread.PC())
        bcr.SetCode(currentFrame.Method().Code())

        
        executeOneInstruction(currentFrame)
    }
}

func executeOneInstruction(frame *rtda.Frame) {
    // todo
    
}
