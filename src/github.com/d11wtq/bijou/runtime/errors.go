package runtime

import (
	"fmt"
)

// Generic runtime error
type RuntimeError struct {
	Message string
}

func (e *RuntimeError) Error() string {
	return e.Message
}

// Error caused by bad function/form arguments
type ArgumentError struct {
	Message string
}

func (e *ArgumentError) Error() string {
	return e.Message
}

// Runtime error caused by impossible airthmetic
type ArithmeticError struct {
	Message string
}

func (e *ArithmeticError) Error() string {
	return e.Message
}

// Runtime error caused by invalid type
type TypeError struct {
	Message string
}

func (e *TypeError) Error() string {
	return e.Message
}

// Return a TypeError for incorrect data types
func BadType(wanted, received Type) error {
	return &TypeError{
		fmt.Sprintf(
			"wrong data type (wanted %s, got %s)",
			TypeName(wanted),
			TypeName(received),
		),
	}
}

// Return an ArgumentError for incorrect arity
func BadArity(wanted, received int) error {
	return &ArgumentError{
		fmt.Sprintf(
			"wrong number of arguments (wanted %d, got %d)",
			wanted,
			received,
		),
	}
}

// Return an ArithmeticError for division by zero
func DivisionByZero() error {
	return &ArithmeticError{"Divide by zero"}
}

// Return a TypeError when calculation is not implemented
func BadOperation(op string, t Type) error {
	return &TypeError{
		fmt.Sprintf(
			"Type %s does not support %s",
			TypeName(t),
			op,
		),
	}
}
