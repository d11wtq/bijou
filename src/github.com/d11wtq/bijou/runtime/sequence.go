package runtime

// Data structures that can be looped over
//
// The semantics of where Head(), Tail() and Put() operate are subject to the
// underlying data structures.
type Sequence interface {
	Value
	// Take the first element of the sequence
	Head() Value
	// Get everything except the first element
	Tail() Sequence
	// Add a new value to the sequence
	Put(Value) (Sequence, error)
	// True if there are no elements in the sequence
	Empty() bool
}
