package runtime

// Macro data type
type Macro Proc

func (mc *Macro) Eq(other Value) bool {
	return mc == other
}

func (mc *Macro) Gt(other Value) bool {
	y, ok := other.(*Macro)
	if ok == false {
		return mc.Type() > other.Type()
	}
	return PtrGt(mc, y)
}

func (mc *Macro) Lt(other Value) bool {
	y, ok := other.(*Macro)
	if ok == false {
		return mc.Type() < other.Type()
	}
	return PtrLt(mc, y)
}

func (mc *Macro) Type() Type {
	return MacroType
}

func (mc *Macro) Eval(env Env) (Value, error) {
	return mc, nil
}

func (mc *Macro) String() string {
	return "#<macro>"
}

// Call this macro with the given arguments
func (mc *Macro) Expand(env Env, args Sequence) (res Value, err error) {
	res, err = (*Proc)(mc).Call(env, args)
	t, ok := res.(TailCall)
	if ok == true {
		res, err = t.Return()
	}

	return
}
