package runtime

// Process the elements of the 'fn' special form
func EvalFn(env Env, s Sequence) (Value, error) {
	if s.Empty() {
		return nil, &RuntimeError{"Missing parameter list in fn"}
	}

	params, err := ValidateParams(s.Head())
	if err != nil {
		return nil, err
	}

	return &Func{
		Params: params,
		Body:   s.Tail(),
		Env:    env,
	}, nil
}

// Ensure the parameter list is syntactically valid
func ValidateParams(s Value) (Sequence, error) {
	params, ok := IsList(s)
	if ok == false {
		return nil, &RuntimeError{"Invalid parameter list type: not a list"}
	}

	p := params
	for !p.Empty() {
		k, ok := p.Head().(Symbol)

		if ok == false {
			return nil, &RuntimeError{"Invalid parameter: not a symbol"}
		}

		if k == Symbol("&") && Length(p) > 2 {
			return nil, &RuntimeError{
				"Positional parameters cannot follow variadic parameters",
			}
		}
		p = p.Tail()
	}

	return params, nil
}
