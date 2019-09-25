package jerrors

type ClassNotFoundError struct {
	name string
}

func NewClassNotFoundError(name string) ClassNotFoundError {
	return ClassNotFoundError{name}
}

func (self ClassNotFoundError) Error() string {
	return self.name
}
