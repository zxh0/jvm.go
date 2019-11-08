package vm

import (
	"fmt"
)

type ErrOrEx struct {
	ClassName string
	Message   string
}

func (err ErrOrEx) Error() string {
	return fmt.Sprintf("%s: %s", err.ClassName, err.Message)
}

func newErrOrEx(className, message string) ErrOrEx {
	return ErrOrEx{
		ClassName: className,
		Message:   message,
	}
}

// LinkageErrors
func NewIncompatibleClassChangeError(msg string) ErrOrEx {
	return newErrOrEx("java/lang/IncompatibleClassChangeError", msg)
}
func NewNoSuchFieldError(msg string) ErrOrEx {
	return newErrOrEx("java/lang/NoSuchFieldError", msg)
}
func NewNoSuchMethodError(msg string) ErrOrEx {
	return newErrOrEx("java/lang/NoSuchMethodError", msg)
}
func NewAbstractMethodError(msg string) ErrOrEx {
	return newErrOrEx("java/lang/AbstractMethodError", msg)
}

func NewClassNotFoundEx(missingClassName string) ErrOrEx {
	return newErrOrEx("java/lang/ClassNotFoundException", missingClassName)
}

func NewNPE() {

}
