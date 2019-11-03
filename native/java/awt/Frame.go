package awt

import (
	"github.com/zxh0/jvm.go/native"
)

func _frame(method native.Method, name, desc string) {
	native.Register("java/awt/Frame", name, desc, method)
}
