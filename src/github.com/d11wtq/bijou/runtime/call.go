package runtime

// Tail calls are handled by a trampoline
type Call struct {
	*NilObj
	// The function that was invoked
	Fn Callable
	// The arguments it was invoked in
	Args Sequence
	// The environment it was invoked in
	Env Env
}

// Resolve the trampoline down to a final return value
func (c *Call) Return() (acc Value, err error) {
	var ok bool

	for {
		acc, err = c.Fn.Call(c.Env, c.Args)
		c, ok = acc.(*Call)
		if ok == false {
			break
		}
	}

	return
}
