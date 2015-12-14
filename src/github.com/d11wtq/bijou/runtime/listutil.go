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
	if ok == false {
		s, ok = v.(*Vector)
	}
	return s, ok
}

func ListEq(list Sequence, other Value) bool {
	if other, ok := IsList(other); ok == true {
		a, b := Sequence(list), Sequence(other)

		for {
			if a.Empty() && b.Empty() {
				return true
			}

			if a.Empty() || b.Empty() {
				return false
			}

			if !Eq(a.Head(), b.Head()) {
				return false
			}

			a, b = a.Tail(), b.Tail()
		}
	} else {
		return false
	}
}

func ListGt(list Sequence, other Value) bool {
	y, ok := IsList(other)
	if ok == false {
		return list.Type() > other.Type()
	}

	a, b := Sequence(list), Sequence(y)

	for {
		if a.Empty() && b.Empty() {
			return false
		}

		if b.Empty() {
			return true
		}

		if a.Empty() {
			return false
		}

		if Gt(a.Head(), b.Head()) {
			return true
		}

		if Lt(a.Head(), b.Head()) {
			return false
		}

		a, b = a.Tail(), b.Tail()
	}
}

func ListLt(list Sequence, other Value) bool {
	y, ok := IsList(other)
	if ok == false {
		return list.Type() < other.Type()
	}

	a, b := Sequence(list), Sequence(y)

	for {
		if a.Empty() && b.Empty() {
			return false
		}

		if a.Empty() {
			return true
		}

		if b.Empty() {
			return false
		}

		if Lt(a.Head(), b.Head()) {
			return true
		}

		if Gt(a.Head(), b.Head()) {
			return false
		}

		a, b = a.Tail(), b.Tail()
	}
}

func ListBind(env Env, list Sequence, value Value) (err error) {
	if v, ok := IsList(value); ok == true {
		var pattern Value

		a, b := Sequence(list), Sequence(v)

		for {
			if a.Empty() && b.Empty() {
				return nil
			}

			if a.Empty() { // pattern empty, but values remain
				return BadPattern(list, value)
			}

			pattern = a.Head()

			if pattern == Symbol("&") {
				if a.Tail().Empty() { // ignored values
					return nil
				}

				a = a.Tail()
				// consume everything into next part of pattern
				pattern, value = a.Head(), b
				b = EmptyList
				if !a.Tail().Empty() {
					return &ArgumentError{"invalid pattern"}
				}
			} else if b.Empty() { // more pattern to match, but no values
				return BadPattern(list, value)
			} else {
				value = b.Head()
			}

			err = Bind(pattern, value, env)
			if err != nil {
				return err
			}

			a, b = a.Tail(), b.Tail()
		}
	} else {
		return BadPattern(list, value)
	}
}
