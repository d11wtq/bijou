package runtime

// Ports are effectively streams of I/O, not limited to characters.
type Port interface {
	Value
	// Write a value to the port. Semantics are port-specific.
	Write(Value) error
	// Take one item from the port. If at EOF, return Nil.
	Accept() (Value, error)
	// Accumulate n units from the port
	Read(n int) (Sequence, error)
	// Close the port so no further I/O can occur
	Close() error
}
