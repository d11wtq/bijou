package runtime

// Predicate to check if a Value is a List or a Cons
//
// Lists and Cons should be considered the same conceptually, thus parts of the
// runtime check for either of the two.
func IsList(v Value) (Sequence, bool) {
	var (
		s  Sequence
		ok bool
	)
	s, ok = v.(*ConsCell)
	if ok == false {
		s, ok = v.(*List)
	}
	return s, ok
}

// Return the length of a given Sequence.
func Length(s Sequence) int {
	acc := 0
	for !s.Empty() {
		acc += 1
		s = s.Tail()
	}
	return acc
}
