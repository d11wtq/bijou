package runtime

import (
	"fmt"
)

// Concrete implementation of Env
type Scope struct {
	// Parent scope, if not the root scope
	Parent Env
	// List of defined symbols in the current scope
	Values map[string]Value
}

// Create a new Scope, extending the parent env.
func NewScope(parent Env) Env {
	return &Scope{
		Parent: parent,
		Values: make(map[string]Value),
	}
}

// Define a new symbol in the current scope
func (s *Scope) Def(k string, v Value) error {
	x, ok := s.Values[k]
	if ok && !x.Eq(v) {
		return &RuntimeError{
			fmt.Sprintf("Attempted to def %s more than once", k),
		}
	}

	s.Values[k] = v
	return nil
}

// Retrieve a symbol from the current scope, or any parent scope
func (s *Scope) Get(k string) (Value, bool) {
	v, ok := s.Values[k]
	if !ok {
		if s.Parent != nil {
			return s.Parent.Get(k)
		}
	}
	return v, ok
}

// Create a new child scope
func (s *Scope) Extend() Env {
	return NewScope(s)
}
