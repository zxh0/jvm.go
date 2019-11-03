package awt

import (
	"github.com/zxh0/jvm.go/native"
)

func init() {
}

func _comp(method native.Method, name, desc string) {
	native.Register("java/awt/Component", name, desc, method)
}
