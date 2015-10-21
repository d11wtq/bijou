package runtime

// List data type
type List struct {
	// Value head at this element of the list
	Data Value
	// The rest of the list (end == nil)
	Next *List
}

// Representation of the empty list
var EmptyList = (*List)(nil)

func (lst *List) Eval(env Env) (Value, error) {
	if lst == nil {
		return lst, nil
	}

	switch lst.Data {
	case Symbol("do"):
		return lst.Next.EvalDo(env)
	}

	return lst, nil
}

func (lst *List) Type() Type {
	return ListType
}

// Get the first element of this list
func (lst *List) Head() Value {
	if lst == nil {
		return Nil
	}

	return lst.Data
}

// Get the next sequence of this list
func (lst *List) Tail() *List {
	if lst == nil {
		return lst
	}

	return lst.Next
}

// Append a new element to the head of the list
func (lst *List) Cons(head Value) *List {
	return &List{head, lst}
}

func (lst *List) EvalDo(env Env) (res Value, err error) {
	for x := lst; x != nil; x = x.Next {
		res, err = x.Data.Eval(env)
		if err != nil {
			return
		}
	}

	return
}
