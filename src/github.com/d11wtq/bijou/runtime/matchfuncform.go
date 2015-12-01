package runtime

// Process the elements of the 'match' function special form
func EvalMatchFunc(env Env, s Sequence) (Value, error) {
	cases := make([]*Func, 0, Length(s))

	for !s.Empty() {
		fn, err := Eval(s.Head(), env)
		if err != nil {
			return nil, err
		}

		if fn, ok := fn.(*Func); ok == true {
			cases = append(cases, fn)
			s = s.Tail()
			continue
		}

		return nil, &ArgumentError{"non-function case in match"}
	}

	return &MatchFunc{cases}, nil
}
