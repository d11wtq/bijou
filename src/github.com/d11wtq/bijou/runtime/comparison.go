package runtime

// Implementation of the = operator.
type Comparable interface {
	Eq(Value) bool
}

// Implementation of the >, < operators.
type Sortable interface {
	Gt(Value) bool
	Lt(Value) bool
}

func Eq(a, b Value) bool {
	if a, ok := a.(Comparable); ok == true {
		return a.Eq(b)
	}

	return a == b
}

func Gt(a, b Value) bool {
	if a, ok := a.(Sortable); ok == true {
		return a.Gt(b)
	}

	if a.Type() == b.Type() {
		return PtrGt(a, b)
	}

	return a.Type() > b.Type()
}

func Lt(a, b Value) bool {
	if a, ok := a.(Sortable); ok == true {
		return a.Lt(b)
	}

	if a.Type() == b.Type() {
		return PtrLt(a, b)
	}

	return a.Type() < b.Type()
}
