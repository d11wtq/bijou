package runtime

// Process the elements of the 'do' special form
func EvalDo(env Env, seq Sequence) (res Value, err error) {
	res = Nil
	for !seq.Empty() {
		// Not a true tail call, resolve now
		if tmp, tco := res.(TailCall); tco {
			res, err = tmp.Resolve()
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
