package runtime

// Value that can be invoked
type Callable interface {
	Value
	// Invoke this value with the given arguments
	Call(env Env, args Sequence) (Value, error)
}
