package awt

import (
	"github.com/zxh0/jvm.go/native"
)

func init() {
}

func _font(method native.Method, name, desc string) {
	native.Register("java/awt/Font", name, desc, method)
}
