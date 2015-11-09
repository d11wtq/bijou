package runtime

// Process the elements of a function call form
func EvalCall(env Env, s Sequence) (Value, error) {
	callee, err := s.Head().Eval(env)
	if err != nil {
		return nil, err
	}

	callable, ok := callee.(Callable)
	if ok == false {
		return nil, &RuntimeError{"Attempted to call non-callable value"}
	}

	args, err := s.Tail(), nil

	if callable.IsMacro() {
		v, err := callable.Call(env, args)
		if err != nil {
			return nil, err
		}
		return v.Eval(env)
	} else {
		args, err = EvalEach(env, args)
		if err != nil {
			return nil, err
		}
		return callable.Call(env, args)
	}
}
