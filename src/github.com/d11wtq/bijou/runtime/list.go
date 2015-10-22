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
	if lst == EmptyList {
		return EmptyList, nil
	}

	switch lst.Data {
	case Symbol("if"):
		return lst.Next.EvalIf(env)
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

// Evaluate every element of this list and return the last value
func (lst *List) EvalDo(env Env) (res Value, err error) {
	for x := lst; x != EmptyList; x = x.Next {
		res, err = x.Data.Eval(env)
		if err != nil {
			return
		}
	}

	return
}

// Evaluate the elements of this list as the 'if' special form
func (lst *List) EvalIf(env Env) (Value, error) {
	if lst == EmptyList {
		return Nil, nil // FIXME: Untested. Should we error?
	}

	if condition, err := lst.Data.Eval(env); err != nil {
		return nil, err
	} else {
		body := lst.Next

		if body != EmptyList {
			if condition != Nil {
				return body.Data.Eval(env)
			} else {
				return body.Next.EvalDo(env)
			}
		}
	}

	return Nil, nil
}
