package test

import (
	"bytes"
	. "github.com/d11wtq/bijou/runtime"
)

// A special scope for use in testing
type FakeScope struct {
	*Scope
}

// Return an implementation of Env, suitable for testing
func FakeEnv() Env {
	scope := &FakeScope{
		&Scope{
			Parent: nil,
			Values: make(map[string]Value),
		},
	}
	scope.Def("*stack-depth*", Int(1))
	return scope
}

// Extend, but set a special var *stack-depth*.
func (s *FakeScope) Extend() Env {
	scope := &FakeScope{
		&Scope{
			Parent: s,
			Values: make(map[string]Value),
		},
	}
	depth, _ := s.Get("*stack-depth*")
	scope.Def("*stack-depth*", Int(1)+depth.(Int))
	return scope
}

// A Fake Value tracking evaluation for testing
type FakeValue struct {
	// The value this fake acts as
	Delegate Value
	// Whether or not this value has been evaluated
	Evaluated bool
	// Result from the delegate
	Result Value
	// Error from the delegate
	Error error
}

// Create a new FakeValue for testing
func NewFakeValue(v Value) *FakeValue {
	return &FakeValue{Delegate: v}
}

func (v *FakeValue) Type() Type {
	return v.Delegate.Type()
}

func (v *FakeValue) Eval(env Env) (Value, error) {
	v.Evaluated = true
	v.Result, v.Error = v.Delegate.Eval(env)
	return v.Result, v.Error
}

func (v *FakeValue) String() string {
	return v.Delegate.String()
}

// A fake function, ignoring args and evaluating tail.
func FakeFn(tail Value) Callable {
	return &Func{
		Env:    FakeEnv(),
		Params: EmptyList.Append(Symbol("&")).Append(Symbol("args")),
		Body:   EmptyList.Append(tail),
	}
}

// Helper to construct lists for testing
func NewList(vs ...Value) *List {
	acc := EmptyList
	for _, v := range vs {
		acc = acc.Append(v)
	}
	return acc
}

// A buffer for fake IO handling
type Buffer struct {
	*bytes.Buffer
	// Flag when the buffer is closed
	Closed bool
}

// Return a new fake IO Buffer.
func FakeIO(buf []byte) *Buffer {
	return &Buffer{bytes.NewBuffer(buf), false}
}

// Set the closed flag on the buffer.
func (buf *Buffer) Close() error {
	buf.Closed = true
	return nil
}
