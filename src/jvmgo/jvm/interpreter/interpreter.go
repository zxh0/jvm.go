package interpreter

import (
    "fmt"
    "jvmgo/jvm/rtda"
    "jvmgo/jvm/instructions"
)

// todo
func Loop(thread *rtda.Thread) {
    defer _debug(thread) // todo

    decoder := instructions.NewInstructionDecoder()
    for {
        frame := thread.CurrentFrame() 

        // decode
        code := frame.Method().Code()
        pc := thread.PC()
        _, inst, nextPC := decoder.Decode(code, pc)
        frame.SetNextPC(nextPC)

        // execute
        //_logInstruction(frame, opcode, inst)
        inst.Execute(frame)
        if !thread.IsStackEmpty() {
            topFrame := thread.TopFrame()
            thread.SetPC(topFrame.NextPC())
        } else {
            break;
        }
    }

    // terminate thread
    threadObj := thread.JThread()
    threadObj.Monitor().NotifyAll()
}

// todo
func _debug(thread *rtda.Thread) {
    if r := recover(); r != nil {
        for !thread.IsStackEmpty() {
            frame := thread.PopFrame()
            fmt.Printf("%v %v\n", frame.Method().Class(), frame.Method())
        }

        err, ok := r.(error)
        if !ok {
            err = fmt.Errorf("%v", r)
            panic(err.Error())
        } else {
            panic(err.Error())
        }
    }
}

func _logInstruction(frame *rtda.Frame, opcode uint8, inst instructions.Instruction) {
    thread := frame.Thread()
    method := frame.Method()
    className := method.Class().Name()
    methodName := method.Name()
    pc := thread.PC()

    if method.IsStatic() {
        fmt.Printf("[instruction] thread:%p %v.%v() #%v 0x%x %v\n", thread, className, methodName, pc, opcode, inst)
    } else {
        fmt.Printf("[instruction] thread:%p %v#%v() #%v 0x%x %v\n", thread, className, methodName, pc, opcode, inst)
    }
}
