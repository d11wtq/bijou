package runtime

// Process the elements of the 'fn' special form
func EvalFn(env Env, lst *List) (Value, error) {
	if lst == EmptyList {
		return nil, &RuntimeError{"Missing parameter list in fn"}
	}
	params, ok := lst.Data.(*List)
	if ok == false {
		return nil, &RuntimeError{"Invalid parameter list type"}
	}

	return &Func{
		Params: params,
		Body:   lst.Next,
		Env:    env,
	}, nil
}
