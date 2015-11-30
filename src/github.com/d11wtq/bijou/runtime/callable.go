package runtime

// Value that can be invoked
type Callable interface {
	Value
	// Call this function with the given arguments at call site
	Call(site Env, args Sequence) (Value, error)
}
