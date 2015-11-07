package runtime

// Read the expected number of arguments into the slice of pointers ptrs
func ReadArgs(args *List, ptrs ...*Value) error {
	for n, ptr := range ptrs {
		if args == EmptyList {
			return BadArity(len(ptrs), n)
		}

		*ptr = args.Data
		args = args.Next
	}

	if args != EmptyList {
		return BadArity(len(ptrs), len(ptrs)+args.Length())
	}

	return nil
}
