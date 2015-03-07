package jerrors

type ClassNotFoundError struct {
	name string
}

func (self ClassNotFoundError) Error() string {
    return self.name
}
