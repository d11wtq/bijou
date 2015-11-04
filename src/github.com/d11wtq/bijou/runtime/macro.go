package runtime

// Macro data type
type Macro struct {
	Func
}

func (mc *Macro) Type() Type {
	return MacroType
}

func (mc *Macro) Eq(other Value) bool {
	return mc == other
}

func (mc *Macro) IsMacro() bool {
	return true
}

func (mc *Macro) Eval(env Env) (Value, error) {
	return mc, nil
}

// Process the elements of the 'macro' special form
func EvalMacro(env Env, lst *List) (Value, error) {
	if lst == EmptyList {
		return nil, &RuntimeError{"Missing parameter list in macro"}
	}
	params, ok := lst.Data.(*List)
	if ok == false {
		return nil, &RuntimeError{"Invalid parameter list type"}
	}

	return &Macro{
		Func{
			Params: params,
			Body:   lst.Next,
			Env:    env,
		},
	}, nil
}
