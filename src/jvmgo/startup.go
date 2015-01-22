package main

import (
    . "jvmgo/any"
    "jvmgo/cmdline"
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

func startJVM(cmd *cmdline.Command) {
    createJVM(cmd)
}

func createJVM(cmd *cmdline.Command) {
    className := cmd.Class()
    classPath := cmd.Options().Classpath()
    classLoader := class.NewClassLoader(classPath)
    createMainThread(className, classLoader)
}

func createMainThread(className string, classLoader Any) {
    //fakeFields := []Any{className}
    //fakeRef := class.NewObj(fakeFields)

    mainThread := rtda.NewThread(128)
    mainThread.PushFrame(createFakeFrame())
}

func createFakeFrame() (*rtda.Frame) {
    //fakeFrame := rtda.NewFrame()
    return nil
}
