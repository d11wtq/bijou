package runtime

// Function data type
type Func Proc

func (fn *Func) Type() Type {
	return FuncType
}

func (fn *Func) Eval(env Env) (Value, error) {
	return fn, nil
}

func (fn *Func) String() string {
	return "#<function>"
}

// Call this function with the given arguments
func (fn *Func) Call(env Env, args Sequence) (Value, error) {
	return (*Proc)(fn).Call(env, args)
}
