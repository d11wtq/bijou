package runtime

// Tail calls are handled by a trampoline
type Trampoline struct {
	*NilObj
	// The function that was invoked
	Fn Callable
	// The arguments it was invoked in
	Args Sequence
	// The environment it was invoked in
	Env Env
}

// Resolve the trampoline down to a final return value
func (t *Trampoline) Resolve() (acc Value, err error) {
	var ok bool

	for {
		acc, err = t.Fn.Call(t.Env, t.Args)
		if err != nil {
			break
		}

		t, ok = acc.(*Trampoline)
		if ok == false {
			break
		}
	}

	return acc, err
}
