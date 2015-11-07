package runtime

// Process the elements of a def form
func EvalDef(env Env, lst *List) (Value, error) {
	if lst == EmptyList {
		return nil, &RuntimeError{"Missing name in `def'"}
	}

	key, ok := lst.Data.(Symbol)
	if ok == false {
		return nil, &RuntimeError{"Bad name in `def' (Symbol required)"}
	}

	if lst.Next == EmptyList {
		return nil, &RuntimeError{"Missing value in `def'"}
	}

	val, err := lst.Next.Data.Eval(env)
	if err != nil {
		return nil, err
	}

	if err := env.Def(string(key), val); err != nil {
		return nil, err
	}

	return val, nil
}
