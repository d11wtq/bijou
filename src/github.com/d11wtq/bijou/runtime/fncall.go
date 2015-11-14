package runtime

// Process the elements of a function call form
func EvalFnCall(env Env, fn Callable, args Sequence) (Value, error) {
	args, err := EvalEach(env, args)
	if err != nil {
		return nil, err
	}

	return &Call{
		Fn:   fn,
		Args: args,
		Env:  env,
	}, nil
}
