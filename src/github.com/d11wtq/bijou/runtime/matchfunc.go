package runtime

// Multiple pattern matching function type
type MatchFunc struct {
	// Function cases to try
	Cases []*Func
}

func (f *MatchFunc) Type() Type {
	return FuncType
}

func (f *MatchFunc) Eval(env Env) (Value, error) {
	return f, nil
}

func (f *MatchFunc) String() string {
	return "#<function>"
}

func (f *MatchFunc) Call(site Env, args Sequence) (Value, error) {
	var (
		env Env
		err error
	)

	for _, fn := range f.Cases {
		env, err = fn.BindEnv(site, args)
		if err == nil {
			return fn.DoBody(env)
		}
	}

	return nil, &PatternError{"none of the patterns matched"}
}
