package runtime

// Process the elements of a quote form
func EvalQuote(env Env, lst *List) (Value, error) {
	return lst.Head(), nil
}
