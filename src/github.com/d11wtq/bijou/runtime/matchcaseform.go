package runtime

// Process the elements of the 'match' case special form
func EvalMatchCase(env Env, value Value, s Sequence) (Value, error) {
	for !s.Empty() {
		env := env.Extend()
		pattern, body := s.Head(), s.Tail().Head()
		err := Bind(pattern, value, env)
		if err == nil {
			return body.Eval(env)
		}
		s = s.Tail().Tail()
	}

	return nil, &PatternError{"none of the cases matched"}
}
