package awt

import (
	"github.com/zxh0/jvm.go/native"
)

func init() {
}

func _tk(method native.Method, name, desc string) {
	native.Register("java/awt/Toolkit", name, desc, method)
}
