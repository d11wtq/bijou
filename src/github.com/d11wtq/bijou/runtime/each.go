package runtime

// Evaluate every element of a list, retaining the list structure
// This is used in argument list evaluation
func EvalEach(env Env, lst *List) (*List, error) {
	acc := EmptyList
	for x := lst; x != EmptyList; x = x.Next {
		v, err := x.Data.Eval(env)
		if err != nil {
			return nil, err
		}
		acc = acc.Cons(v)
	}
	return acc.Reverse(), nil
}
