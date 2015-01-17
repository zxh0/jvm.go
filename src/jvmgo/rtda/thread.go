package rtda

/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/

type Thread struct {
    // todo
    pc      uint32
    stack   *Stack
}

type Stack struct {
    frames  []*Frame
}
