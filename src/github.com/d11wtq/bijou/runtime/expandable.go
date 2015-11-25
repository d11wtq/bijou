package runtime

// Value that can be expanded with arguments
type Expandable interface {
	Value
	// Transform this value into another syntactic element
	Expand(env Env, args Sequence) (Value, error)
}
