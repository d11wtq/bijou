package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestBooleanType(t *testing.T) {
	if True.Type() != BooleanType {
		t.Fatalf(`expected True.Type() == BooleanType, got %s`, True.Type())
	}

	if False.Type() != BooleanType {
		t.Fatalf(`expected False.Type() == BooleanType, got %s`, False.Type())
	}
}

func TestBooleanEvalToSelf(t *testing.T) {
	env := FakeEnv()
	v1, err := True.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v1 != True {
		t.Fatalf(`expected v1 == True, got %s`, v1)
	}

	v2, err := False.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v2 != False {
		t.Fatalf(`expected v2 == False, got %s`, v2)
	}
}

func TestBooleanEq(t *testing.T) {
	if !True.Eq(True) {
		t.Fatalf(`expected True.Eq(True), got false`)
	}
	if !False.Eq(False) {
		t.Fatalf(`expected False.Eq(False), got false`)
	}

	if False.Eq(True) {
		t.Fatalf(`expected !False.Eq(True), got true`)
	}
	if True.Eq(False) {
		t.Fatalf(`expected !True.Eq(False), got true`)
	}
}
