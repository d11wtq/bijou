package runtime

// Generic runtime error
type RuntimeError struct {
	// Contextual error message (user-friendly)
	Message string
}

func (e *RuntimeError) Error() string {
	return e.Message
}
