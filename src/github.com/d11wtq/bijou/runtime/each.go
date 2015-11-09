package runtime

// Evaluate every element of a list, retaining the list structure
// This is used in argument list evaluation
func EvalEach(env Env, s Sequence) (Sequence, error) {
	acc := &List{}
	for !s.Empty() {
		v, err := s.Head().Eval(env)
		if err != nil {
			return nil, err
		}
		acc.Append(v)
		s = s.Tail()
	}
	return acc, nil
}
