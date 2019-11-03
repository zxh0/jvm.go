package awt

import (
	"github.com/zxh0/jvm.go/native"
)

func init() {
}

func _cursor(method native.Method, name, desc string) {
	native.Register("java/awt/Cursor", name, desc, method)
}
