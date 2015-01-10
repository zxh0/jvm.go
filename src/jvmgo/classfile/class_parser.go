package classfile

//import "errors"
import "fmt"

func ParseClassFile(reader *ClassReader) (cf *ClassFile, err error) {
    defer func() {
        if r := recover(); r != nil {
            var ok bool
            err, ok = r.(error)
            if !ok {
                err = fmt.Errorf("%v", r)
            }
        }
    }()

    cf = &ClassFile{}
    readClass(cf, reader)
    return
}
