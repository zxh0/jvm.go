package heap

var EmptySlot = Slot{0, nil}

type Slot struct {
	Val int64 // big enough to hold any primitive value
	Ref *Object
}

func NewRefSlot(ref *Object) Slot {
	return Slot{0, ref}
}
