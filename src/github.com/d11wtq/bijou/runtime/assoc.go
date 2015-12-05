package runtime

// Associative data types, such as maps and modules
type Associative interface {
	// Lookup a value by key, or return an error
	Lookup(Value) (Value, error)
}
