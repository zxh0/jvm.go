package main

import (
    . "jvmgo/any"
    "jvmgo/cmdline"
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

func startJVM(cmd *cmdline.Command) {
    className := cmd.Class()
    classPath := cmd.Options().Classpath()
    classLoader := class.NewClassLoader(classPath)
    mainThread := createMainThread(className, classLoader)
    loop(mainThread)
}

func createMainThread(className string, classLoader Any) (*rtda.Thread) {
    fakeFields := []Any{className, classLoader}
    fakeRef := class.NewObj(fakeFields)

    fakeMethod := class.NewStartupMethod([]byte{0xff})
    fakeFrame := rtda.NewFrame(fakeMethod)
    fakeFrame.OperandStack().PushRef(fakeRef)

    mainThread := rtda.NewThread(128)
    mainThread.PushFrame(fakeFrame)
    return mainThread
}
