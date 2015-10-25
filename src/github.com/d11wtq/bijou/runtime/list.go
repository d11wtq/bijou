package runtime

// List data type (Seq)
type List struct {
	// Value head at this element of the list
	Data Value
	// The rest of the list (end == nil)
	Next *List
}

// Representation of the empty list
var EmptyList = (*List)(nil)

func (lst *List) Eval(env Env) (Value, error) {
	if lst == EmptyList {
		return EmptyList, nil
	}

	switch lst.Data {
	case Symbol("if"):
		return EvalIf(env, lst.Next)
	case Symbol("do"):
		return EvalDo(env, lst.Next)
	case Symbol("fn"):
		return EvalFn(env, lst.Next)
	default:
		return EvalCall(env, lst)
	}
}

func (lst *List) Type() Type {
	return ListType
}

// Get the first element of this list
func (lst *List) Head() Value {
	if lst == EmptyList {
		return Nil
	}

	return lst.Data
}

// Get the next sequence of this list
func (lst *List) Tail() *List {
	if lst == EmptyList {
		return EmptyList
	}

	return lst.Next
}

// Append a new element to the head of the list
func (lst *List) Cons(head Value) *List {
	return &List{head, lst}
}
