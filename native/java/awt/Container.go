package awt

import (
	"github.com/zxh0/jvm.go/native"
)

func init() {
}

func _container(method native.Method, name, desc string) {
	native.Register("java/awt/Container", name, desc, method)
}
