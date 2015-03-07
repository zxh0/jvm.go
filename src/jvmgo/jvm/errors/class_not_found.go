package exception

type ClassNotFoundError struct {
	name string
}

func (self ClassNotFoundError) Error() string {
    return "class not found: " + self.name
}
