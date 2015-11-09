package runtime

// List data type (Sequence)
type List struct {
	// Otherwise behaves like a ConsCell
	*ConsCell
	// Retain a pointer to the last element in the list
	Last *ConsCell
}

func (lst *List) Type() Type {
	return ListType
}

func (lst *List) Eval(env Env) (Value, error) {
	if lst.Empty() {
		return lst, nil
	}

	return lst.ConsCell.Eval(env)
}

func (lst *List) Put(v Value) (Sequence, error) {
	newLst := &List{&ConsCell{v, lst}, lst.Last}
	if newLst.Last.Empty() {
		newLst.Last = newLst.ConsCell
	}
	return newLst, nil
}

// Append a value to this List at the end.
//
// WARNING: This method is destructive and only intended to be used during list
//          construction. Attempts to use it at runtime will violate the
//          immutability principle.
func (lst *List) Append(v Value) *List {
	newCons := &ConsCell{v, EmptyCons}
	if lst.Last.Empty() {
		lst.ConsCell = newCons
		lst.Last = newCons
		return lst
	}

	lst.Last.Next = newCons
	lst.Last = newCons
	return lst
}
