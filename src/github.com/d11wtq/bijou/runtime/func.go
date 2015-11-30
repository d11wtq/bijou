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

func (fn *Func) BindEnv(site Env, args Sequence) (Env, error) {
	return (*Proc)(fn).BindEnv(site, args)
}

func (fn *Func) DoBody(env Env) (Value, error) {
	return (*Proc)(fn).DoBody(env)
}

func (fn *Func) Call(site Env, args Sequence) (Value, error) {
	return (*Proc)(fn).Call(site, args)
}
