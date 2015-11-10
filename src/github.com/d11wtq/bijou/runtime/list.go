package runtime

// List data type (Sequence)
type List struct {
	// Otherwise behaves like a ConsCell
	*ConsCell
	// Retain a pointer to the last element in the list
	Last *List
}

var EmptyList = &List{}

func (lst *List) Type() Type {
	return ListType
}

func (lst *List) Eval(env Env) (Value, error) {
	if lst.Empty() {
		return EmptyList, nil
	}

	return lst.ConsCell.Eval(env)
}

func (lst *List) Tail() Sequence {
	if lst.Empty() {
		return EmptyList
	}

	return lst.Next
}

func (lst *List) Put(v Value) (Sequence, error) {
	newLst := &List{&ConsCell{v, lst}, EmptyList}

	if lst.Empty() {
		newLst.Last = newLst
	} else {
		newLst.Last = lst.Last
	}

	return newLst, nil
}

// Append a value to the end of this List (destructive).
func (lst *List) Append(v Value) *List {
	if lst.Empty() {
		lst2, _ := lst.Put(v)
		return lst2.(*List)
	}

	newLst := &List{&ConsCell{v, EmptyList}, EmptyList}
	newLst.Last = newLst

	lst.Last.Next = newLst
	lst.Last = newLst
	return lst
}
