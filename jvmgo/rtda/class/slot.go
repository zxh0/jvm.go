package class

var EmptySlot = Slot{0, nil}

type Slot struct {
	Val int64 // big enough to hold any primitive value
	Ref *Obj
}

func NewRefSlot(ref *Obj) Slot {
	return Slot{0, ref}
}
