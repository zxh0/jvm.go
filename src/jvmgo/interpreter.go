package main

import "jvmgo/rtda"

func interpret() {

}

// todo
func loop(thread *rtda.Thread) {
    for !thread.IsStackEmpty() {
        currentFrame := thread.CurrentFrame()
        executeOneInstruction(currentFrame)
    }
}

func executeOneInstruction(self *rtda.Frame) {
    // todo
}
