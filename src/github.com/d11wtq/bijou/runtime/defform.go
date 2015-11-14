package runtime

// Process the elements of a def form
func EvalDef(env Env, s Sequence) (Value, error) {
	var key, doc, val Value

	if ReadArgs(s, &key, &doc, &val) != nil {
		doc = String("")
		err := ReadArgs(s, &key, &val)
		if err != nil {
			switch {
			case (key == nil):
				return nil, &ArgumentError{"Missing name in def"}
			case (val == nil):
				return nil, &ArgumentError{"Missing value in def"}
			default:
				return nil, &ArgumentError{"Too many arguments to def"}
			}
		}
	}

	sym, ok := key.(Symbol)
	if ok == false {
		return nil, &ArgumentError{"Bad name in def (symbol required)"}
	}

	val, err := Eval(val, env)
	if err != nil {
		return nil, err
	}

	err = env.Def(string(sym), val)
	if err != nil {
		return nil, err
	}

	return val, nil
}
