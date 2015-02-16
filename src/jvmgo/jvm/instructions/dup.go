package instructions

import (
	"jvmgo/any"
	"jvmgo/jvm/rtda"
)

// Duplicate the top operand stack value
type dup struct{ NoOperandsInstruction }

func (self *dup) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.Pop()
	stack.Push(val)
	stack.Push(val)
}

// Duplicate the top operand stack value and insert two values down
type dup_x1 struct{ NoOperandsInstruction }

func (self *dup_x1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.Pop()
	val2 := stack.Pop()
	stack.Push(val1)
	stack.Push(val2)
	stack.Push(val1)
}

// Duplicate the top operand stack value and insert two or three values down
type dup_x2 struct{ NoOperandsInstruction }

func (self *dup_x2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.Pop()
	val2 := stack.Pop()
	if any.IsLongOrDouble(val2) {
		// form2
		stack.Push(val1)
		stack.Push(val2)
		stack.Push(val1)
	} else {
		// form1
		val3 := stack.Pop()
		stack.Push(val1)
		stack.Push(val3)
		stack.Push(val2)
		stack.Push(val1)
	}
}

// Duplicate the top one or two operand stack values
type dup2 struct{ NoOperandsInstruction }

func (self *dup2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.Pop()
	if any.IsLongOrDouble(val1) {
		// form2
		stack.Push(val1)
		stack.Push(val1)
	} else {
		// form1
		val2 := stack.Pop()
		stack.Push(val2)
		stack.Push(val1)
		stack.Push(val2)
		stack.Push(val1)
	}
}

// Duplicate the top one or two operand stack values and insert two or three values down
type dup2_x1 struct{ NoOperandsInstruction }

func (self *dup2_x1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.Pop()
	if any.IsLongOrDouble(val1) {
		// form2
		val2 := stack.Pop()
		stack.Push(val1)
		stack.Push(val2)
		stack.Push(val1)
	} else {
		// form1
		val2 := stack.Pop()
		val3 := stack.Pop()
		stack.Push(val2)
		stack.Push(val1)
		stack.Push(val3)
		stack.Push(val2)
		stack.Push(val1)
	}
}

// Duplicate the top one or two operand stack values and insert two, three, or four values down
type dup2_x2 struct{ NoOperandsInstruction }

func (self *dup2_x2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.Pop()
	val2 := stack.Pop()
	if any.IsLongOrDouble(val1) {
		if any.IsLongOrDouble(val2) {
			// form4
			//.., value2, value1 →
			//..., value1, value2, value1
			stack.Push(val1)
			stack.Push(val2)
			stack.Push(val1)
		} else {
			// form2
			//..., value3, value2, value1 →
			//..., value1, value3, value2, value1
			val3 := stack.Pop()
			stack.Push(val1)
			stack.Push(val3)
			stack.Push(val2)
			stack.Push(val1)
		}
	} else {
		val3 := stack.Pop()
		if any.IsLongOrDouble(val3) {
			// form3
			//..., value3, value2, value1 →
			//..., value2, value1, value3, value2, value1
			stack.Push(val2)
			stack.Push(val1)
			stack.Push(val3)
			stack.Push(val2)
			stack.Push(val1)
		} else {
			// form1
			//..., value4, value3, value2, value1 →
			//..., value2, value1, value4, value3, value2, value1
			val4 := stack.Pop()
			stack.Push(val2)
			stack.Push(val1)
			stack.Push(val4)
			stack.Push(val3)
			stack.Push(val2)
			stack.Push(val1)
		}
	}
}
