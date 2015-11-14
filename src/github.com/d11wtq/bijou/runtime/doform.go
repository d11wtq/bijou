package runtime

// Process the elements of the 'do' special form
func EvalDo(env Env, seq Sequence) (res Value, err error) {
	res = Nil
	for !seq.Empty() {
		// Not a tail call in this position
		t, ok := res.(TailCall)
		if ok == true {
			res, err = t.Return()
			if err != nil {
				return
			}
		}

		res, err = seq.Head().Eval(env)
		if err != nil {
			return
		}
		seq = seq.Tail()
	}

	return
}
