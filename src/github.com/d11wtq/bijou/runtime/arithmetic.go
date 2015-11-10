package runtime

// Interface for values supporting addition
type Addition interface {
	Value
	// Handle (+ a b) and return a new value
	Add(v Value) (Value, error)
}

// Interface for values supporting subtraction
type Subtraction interface {
	Value
	// Handle (- a b) and return a new value
	Sub(v Value) (Value, error)
}
