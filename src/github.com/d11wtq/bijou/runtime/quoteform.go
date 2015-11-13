package runtime

// Process the elements of a quote form
func EvalQuote(env Env, s Sequence) (Value, error) {
	return s.Head(), nil
}
