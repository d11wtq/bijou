package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
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
	env := test.FakeEnv()
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

func TestBooleanGt(t *testing.T) {
	if True.Gt(True) {
		t.Fatalf(`expected !True.Gt(True), got true`)
	}
	if !True.Gt(False) {
		t.Fatalf(`expected True.Gt(False), got false`)
	}
	if False.Gt(False) {
		t.Fatalf(`expected !False.Gt(False), got true`)
	}
	if False.Gt(True) {
		t.Fatalf(`expected !False.Gt(True), got true`)
	}
	if False.Gt(Int(0)) {
		t.Fatalf(`expected !False.Gt(Int(0)), got true`)
	}
	if !False.Gt(Nil) {
		t.Fatalf(`expected False.Gt(Nil), got false`)
	}
}

func TestBooleanLt(t *testing.T) {
	if True.Lt(True) {
		t.Fatalf(`expected !True.Lt(True), got true`)
	}
	if !False.Lt(True) {
		t.Fatalf(`expected False.Gt(True), got false`)
	}
	if False.Lt(False) {
		t.Fatalf(`expected !False.Gt(False), got true`)
	}
	if True.Lt(False) {
		t.Fatalf(`expected !True.Lt(False), got true`)
	}
	if !False.Lt(Int(0)) {
		t.Fatalf(`expected False.Lt(Int(0)), got false`)
	}
	if !True.Lt(Int(0)) {
		t.Fatalf(`expected True.Lt(Int(0)), got false`)
	}
	if False.Lt(Nil) {
		t.Fatalf(`expected !False.Lt(Nil), got true`)
	}
}

func TestBooleanString(t *testing.T) {
	if True.String() != "true" {
		t.Fatalf(`expected True.String() == "true", got %s`, True.String())
	}

	if False.String() != "false" {
		t.Fatalf(`expected False.String() == "true", got %s`, False.String())
	}
}
