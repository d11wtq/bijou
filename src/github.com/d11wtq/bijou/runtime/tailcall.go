package runtime

// Function calls are emitted as deferred procedures
type TailCall interface {
	// Resolve the tail call to its value
	Return() (Value, error)
}
