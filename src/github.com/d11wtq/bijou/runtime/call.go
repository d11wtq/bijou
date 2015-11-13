package runtime

// Process the elements of a function call form
func EvalCall(env Env, fn Callable, args Sequence) (Value, error) {
	args, err := EvalEach(env, args)
	if err != nil {
		return nil, err
	}
	return fn.Call(env, args)
}
