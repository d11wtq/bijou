package runtime

// Process the elements of a function call form
func EvalCall(env Env, lst *List) (Value, error) {
	callee, err := lst.Data.Eval(env)
	if err != nil {
		return nil, err
	}

	callable, ok := callee.(Callable)
	if ok == false {
		return nil, &RuntimeError{"Attempted to call non-callable value"}
	}

	args, err := lst.Next, nil

	if callable.IsMacro() {
		v, err := callable.Call(args)
		if err != nil {
			return nil, err
		}
		return v.Eval(env)
	} else {
		args, err = EvalEach(env, lst.Next)
		if err != nil {
			return nil, err
		}
		return callable.Call(args)
	}
}
