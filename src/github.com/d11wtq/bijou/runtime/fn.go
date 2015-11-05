package runtime

// Process the elements of the 'fn' special form
func EvalFn(env Env, lst *List) (Value, error) {
	if lst == EmptyList {
		return nil, &RuntimeError{"Missing parameter list in fn"}
	}

	params, err := ValidateParams(lst.Data)
	if err != nil {
		return nil, err
	}

	return &Func{
		Params: params,
		Body:   lst.Next,
		Env:    env,
	}, nil
}

// Ensure the parameter list is syntactically valid
func ValidateParams(seq Value) (*List, error) {
	params, ok := seq.(*List)
	if ok == false {
		return nil, &RuntimeError{"Invalid parameter list type: not a list"}
	}

	for p := params; p != EmptyList; p = p.Next {
		k, ok := p.Data.(Symbol)

		if ok == false {
			return nil, &RuntimeError{"Invalid parameter: not a symbol"}
		}

		if k == Symbol("&") && p.Length() > 2 {
			return nil, &RuntimeError{
				"Positional parameters cannot follow variadic parameters",
			}
		}
	}

	return params, nil
}
