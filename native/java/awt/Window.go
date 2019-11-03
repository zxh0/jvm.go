package awt

import (
	"github.com/zxh0/jvm.go/native"
)

func _window(method native.Method, name, desc string) {
	native.Register("java/awt/Window", name, desc, method)
}
