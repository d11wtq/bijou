package test

import (
	. "github.com/d11wtq/bijou/runtime"
)

// Return an implementation of Env, suitable for testing
func FakeEnv() Env {
	return NewScope(nil)
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

func (v *FakeValue) Eq(other Value) bool {
	return v.Delegate.Eq(other)
}

func (v *FakeValue) Gt(other Value) bool {
	return v.Delegate.Gt(other)
}

func (v *FakeValue) Lt(other Value) bool {
	return v.Delegate.Lt(other)
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

// Helper to construct lists for testing
func NewList(vs ...Value) *List {
	acc := EmptyList
	for _, v := range vs {
		acc = acc.Append(v)
	}
	return acc
}
