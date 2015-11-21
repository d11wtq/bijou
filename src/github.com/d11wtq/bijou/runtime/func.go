package runtime

// Function data type
type Func Proc

func (fn *Func) Eq(other Value) bool {
	return fn == other
}

func (fn *Func) Gt(other Value) bool {
	y, ok := other.(*Func)
	if ok == false {
		return fn.Type() > other.Type()
	}
	return PtrGt(fn, y)
}

func (fn *Func) Lt(other Value) bool {
	y, ok := other.(*Func)
	if ok == false {
		return fn.Type() < other.Type()
	}
	return PtrLt(fn, y)
}

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
