package runtime

// Process the elements of the 'macro' special form
func EvalMacro(env Env, s Sequence) (Value, error) {
	if s.Empty() {
		return nil, &ArgumentError{"Missing parameter list in macro"}
	}

	params, err := ValidateParams(s.Head())
	if err != nil {
		return nil, err
	}

	return &Macro{
		Params: params,
		Body:   s.Tail(),
		Env:    env,
	}, nil
}
