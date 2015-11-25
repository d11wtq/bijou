package runtime

// Interface for values that support binding
type Binding interface {
	// Try to bind the value in Env
	Bind(Env, Value) error
}

// Perform a pattern match between pattern and value in env.
// This has the side-effect of defining variables in env.
// If a match can't be made, returns a PatternError.
func Bind(pattern, value Value, env Env) error {
	binding, ok := pattern.(Binding)
	if ok == true {
		return binding.Bind(env, value)
	}

	if !Eq(pattern, value) {
		return BadPattern(pattern, value)
	}

	return nil
}
