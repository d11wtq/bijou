package runtime

// Read the expected number of arguments into the slice of pointers ptrs
func ReadArgs(args Sequence, ptrs ...*Value) error {
	for n, ptr := range ptrs {
		if args.Empty() {
			return BadArity(len(ptrs), n)
		}

		*ptr = args.Head()
		args = args.Tail()
	}

	if !args.Empty() {
		return BadArity(len(ptrs), len(ptrs)+Length(args))
	}

	return nil
}
