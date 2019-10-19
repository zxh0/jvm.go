package vm

type ClassNotFoundError struct {
	name string
}

func NewClassNotFoundError(name string) ClassNotFoundError {
	return ClassNotFoundError{name}
}

func (err ClassNotFoundError) Error() string {
	return err.name
}
