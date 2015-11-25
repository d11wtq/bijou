package runtime

// Process the elements of a bind form
func EvalBind(env Env, s Sequence) (Value, error) {
	var err error
	var pattern, value Value

	err = ReadArgs(s, &pattern, &value)
	if err != nil {
		if s.Empty() {
			return nil, &ArgumentError{"Missing pattern in bind"}
		}

		if s.Tail().Empty() {
			return nil, &ArgumentError{"Missing value in bind"}
		}

		return nil, &ArgumentError{"Too many arguments to bind"}
	}

	value, err = value.Eval(env)
	if err != nil {
		return nil, err
	}

	err = env.Bind(pattern, value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
