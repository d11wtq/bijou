package runtime

// Invoke function f with args and resolve any tail calls.
func Apply(f Callable, env Env, args Sequence) (res Value, err error) {
	res, err = f.Call(env, args)
	t, ok := res.(TailCall)
	if ok == true {
		res, err = t.Return()
	}

	return
}
