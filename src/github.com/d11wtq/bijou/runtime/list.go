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

func (lst *List) Eq(other Value) bool {
	if other, ok := other.(*List); ok == true {
		for {
			if lst == EmptyList && other == EmptyList {
				return true
			}

			if lst == EmptyList || other == EmptyList {
				return false
			}

			if !lst.Data.Eq(other.Data) {
				return false
			}

			lst, other = lst.Next, other.Next
		}
	} else {
		return false
	}
}

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
	case Symbol("def"):
		return EvalDef(env, lst.Next)
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

// Return a new list with the elements of this list in reverse
func (lst *List) Reverse() *List {
	acc := EmptyList
	for x := lst; x != EmptyList; x = x.Next {
		acc = acc.Cons(x.Data)
	}
	return acc
}
