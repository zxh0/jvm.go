package jerrors

import (
	"fmt"
)

type ClassNotFoundError struct {
	name string
}

func NewClassNotFoundError(name string) ClassNotFoundError {
	return ClassNotFoundError{name}
}

func (self ClassNotFoundError) Error() string {
	return fmt.Sprintf("Class not found %v", self.name)
}
