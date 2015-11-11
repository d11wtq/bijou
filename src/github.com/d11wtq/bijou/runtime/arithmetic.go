package runtime

// Interface for values supporting addition
type Addition interface {
	// Handle (+ a b)
	Add(v Value) (Value, error)
}

// Interface for values supporting subtraction
type Subtraction interface {
	// Handle (- a b)
	Sub(v Value) (Value, error)
}

// Interface for values supporting multiplication
type Multiplication interface {
	// Handle (* a b)
	Mul(v Value) (Value, error)
}

// Interface for values supporting division
type Division interface {
	// Handle (/ a b)
	Div(v Value) (Value, error)
}
