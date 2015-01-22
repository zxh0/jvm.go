package main

import (
    //. "jvmgo/any"
    "jvmgo/cmdline"
    "jvmgo/rtda"
    //"jvmgo/rtda/class"
)

func startJVM(cmd *cmdline.Command) {

}

func createJVM(cmd *cmdline.Command) {
    //classPath := cmd.Options().Classpath()
    className := cmd.Class()

    createMainThread(className)
}

func createMainThread(className string) {
    //fakeFields := []Any{className}
    //fakeRef := class.NewObj(fakeFields)

    mainThread := rtda.NewThread(128)
    mainThread.PushFrame(createFakeFrame())
}

func createFakeFrame() (*rtda.Frame) {
    //fakeFrame := rtda.NewFrame()
    return nil
}
