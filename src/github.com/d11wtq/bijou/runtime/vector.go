package runtime

import (
	"fmt"
	"github.com/d11wtq/persistent/vector"
)

// Persistent vector based on github.com/d11wtq/persistent/vector
type Vector struct {
	// The underlying vector
	Vec *vector.Vector
	// The offset to start at (obtaining a tail)
	Offset uint32
}

// Representation of the empty vector
var EmptyVector = &Vector{
	Vec:    vector.New(),
	Offset: 0,
}

func (vec *Vector) Type() Type {
	return VectorType
}

func (vec *Vector) Eval(env Env) (Value, error) {
	return vec, nil
}

func (vec *Vector) String() string {
	return "#<vector>"
}

func (vec *Vector) Head() Value {
	v, err := vec.Lookup(Int(0))
	if err != nil {
		return Nil
	}

	return v
}

func (vec *Vector) Tail() Sequence {
	if vec.Empty() {
		return vec
	}

	return &Vector{
		Vec:    vec.Vec,
		Offset: vec.Offset + 1,
	}
}

func (vec *Vector) Put(v Value) (Sequence, error) {
	return vec.Append(v), nil
}

func (vec *Vector) Empty() bool {
	return vec.Vec.Count() == vec.Offset
}

func (vec *Vector) Eq(other Value) bool {
	return ListEq(vec, other)
}

func (vec *Vector) Gt(other Value) bool {
	return ListGt(vec, other)
}

func (vec *Vector) Lt(other Value) bool {
	return ListLt(vec, other)
}

func (vec *Vector) Bind(env Env, value Value) error {
	return ListBind(env, vec, value)
}

func (vec *Vector) Lookup(key Value) (Value, error) {
	idx, ok := key.(Int)
	if ok == false {
		return nil, BadType(IntType, key.Type())
	}

	v, err := vec.Vec.Get(uint32(idx) + vec.Offset)
	if err != nil {
		return nil, &RuntimeError{fmt.Sprintf("key %d out of bounds", key)}
	}

	return v.(Value), nil
}

// Append a value to the end of this vector, returning a new vector.
func (vec *Vector) Append(v Value) *Vector {
	return &Vector{
		Vec:    vec.Vec.Append(v),
		Offset: vec.Offset,
	}
}
