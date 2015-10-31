package runtime

import (
	"fmt"
)

// Generic runtime error
type RuntimeError struct {
	// Contextual error message (user-friendly)
	Message string
}

func (e *RuntimeError) Error() string {
	return e.Message
}

// Runtime error caused by bad function arguments
type ArgumentError struct {
	// Contextual error message (user-friendly)
	Message string
}

func (e *ArgumentError) Error() string {
	return e.Message
}

func BadArity(wanted, received int) error {
	return &ArgumentError{
		fmt.Sprintf(
			"wrong number of arguments (wanted %d, got %d)",
			wanted,
			received,
		),
	}
}
