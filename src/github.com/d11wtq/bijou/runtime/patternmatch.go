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
	tx := env.Transaction()

	binding, ok := pattern.(Binding)
	if ok == true {
		err := binding.Bind(tx, value)
		if err != nil {
			return err
		}
		tx.Commit()
		return nil
	}

	if !Eq(pattern, value) {
		return BadPattern(pattern, value)
	}

	return nil
}
