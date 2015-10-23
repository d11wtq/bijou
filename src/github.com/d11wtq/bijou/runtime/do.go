package runtime

// Process the elements of the 'do' special form
func EvalDo(env Env, lst *List) (res Value, err error) {
	res = Nil
	for x := lst; x != EmptyList; x = x.Next {
		res, err = x.Data.Eval(env)
		if err != nil {
			return
		}
	}

	return
}
