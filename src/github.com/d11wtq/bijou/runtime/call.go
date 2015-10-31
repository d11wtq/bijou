package runtime

// Process the elements of a function call form
func EvalCall(env Env, lst *List) (Value, error) {
	callee, err := lst.Data.Eval(env)
	if err != nil {
		return nil, err
	}

	fn, ok := callee.(Callable)
	if ok == false {
		return nil, &RuntimeError{"Attempted to call non-function"}
	}

	args, err := EvalEach(env, lst.Next)
	if err != nil {
		return nil, err
	}

	return fn.Call(args)
}
