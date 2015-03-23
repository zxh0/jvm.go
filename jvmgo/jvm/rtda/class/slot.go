package class

type Slot struct {
	Val int64 // big enough to hold any primitive value
	Ref *Obj
}
