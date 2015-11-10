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

func (mc *Macro) String() string {
	return "#<macro>"
}

// Process the elements of the 'macro' special form
func EvalMacro(env Env, s Sequence) (Value, error) {
	if s.Empty() {
		return nil, &RuntimeError{"Missing parameter list in macro"}
	}
	params, ok := IsList(s.Head())
	if ok == false {
		return nil, &RuntimeError{"Invalid parameter list type"}
	}

	return &Macro{
		Func{
			Params: params,
			Body:   s.Tail(),
			Env:    env,
		},
	}, nil
}
