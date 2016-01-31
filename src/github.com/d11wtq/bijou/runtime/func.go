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
	if fn.Name == "" {
		return (*Proc)(fn).BindEnv(site, args)
	}

	env, err := (*Proc)(fn).BindEnv(site, args)
	if err != nil {
		return env, err
	}
	err = env.Def(fn.Name, fn)
	if err != nil {
		return nil, err
	}
	return env, nil
}

func (fn *Func) DoBody(env Env) (Value, error) {
	return (*Proc)(fn).DoBody(env)
}

func (fn *Func) Call(site Env, args Sequence) (Value, error) {
	env, err := fn.BindEnv(site, args)
	if err != nil {
		return nil, err
	}
	return fn.DoBody(env)
}
