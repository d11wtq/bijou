package runtime

// Program data type
type Type uint8

const (
	// Empty value, nil
	NilType Type = iota
	// Boolean true/false
	BooleanType
	// Integers
	IntType
	// Unique symbols; variable names, function names
	SymbolType
	// Generic sequences
	SequenceType
	// Strings
	StringType
	// Cons sequences
	ConsType
	// Linked lists
	ListType
	// Functions
	FuncType
	// Macros
	MacroType
	// Input/output ports
	PortType
)

// Program runtime value
type Value interface {
	// Eval to the lowest level value
	Eval(env Env) (Value, error)
	// Eval to the lowest level value
	Type() Type
	// Provide a representation of this value as string
	String() string
	// = comparison
	Eq(Value) bool
	// > comparison
	Gt(Value) bool
	// < comparison
	Lt(Value) bool
}

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

// Runtime environment
type Env interface {
	// Define a new symbol
	Def(string, Value) error
	// Lookup a variable
	Get(string) (Value, bool)
	// Introduce a new scope
	Extend() Env
}

// Value that can be invoked
type Callable interface {
	Value
	// Invoke this value with the given arguments
	Call(env Env, args Sequence) (Value, error)
}

// Value that can be expanded with arguments
type Expandable interface {
	Value
	// Transform this value into another syntactic element
	Expand(env Env, args Sequence) (Value, error)
}

// Function calls are emitted as deferred procedures
type TailCall interface {
	// Resolve the tail call to its value
	Return() (Value, error)
}

// Ports are effectively streams of I/O, not limited to characters.
type Port interface {
	Value
	// Write a value to the port. Semantics are port-specific.
	Write(Value) error
	// Take one item from the port. If at EOF, return Nil.
	Accept() (Value, error)
	// Accumulate n units from the port
	Read(n int) (Sequence, error)
	// Close the port so no further I/O can occur
	Close() error
}
