package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestNilType(t *testing.T) {
	if Nil.Type() != NilType {
		t.Fatalf(`expected Nil.Type() == NilType, got %s`, Nil.Type())
	}
}

func TestNilEvalToSelf(t *testing.T) {
	env := FakeEnv()
	v, err := Nil.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != Nil {
		t.Fatalf(`expected v == Nil, got %s`, v)
	}
}

func TestNilEq(t *testing.T) {
	if !Nil.Eq(Nil) {
		t.Fatalf(`expected Nil.Eq(Nil), got false`)
	}
	if Nil.Eq(Int(42)) {
		t.Fatalf(`expected !Nil.Eq(Int(42)), got true`)
	}
}

func TestNilGt(t *testing.T) {
	if Nil.Gt(Nil) {
		t.Fatalf(`expected !Nil.Gt(Nil), got true`)
	}
	if Nil.Gt(False) {
		t.Fatalf(`expected !Nil.Gt(False), got true`)
	}
	if Nil.Gt(Int(0)) {
		t.Fatalf(`expected !Nil.Gt(Int(0)), got true`)
	}
}

func TestNilLt(t *testing.T) {
	if Nil.Lt(Nil) {
		t.Fatalf(`expected !Nil.Lt(Nil), got true`)
	}
	if !Nil.Lt(False) {
		t.Fatalf(`expected Nil.Lt(False), got false`)
	}
	if !Nil.Lt(Int(0)) {
		t.Fatalf(`expected Nil.Lt(Int(0)), got false`)
	}
}

func TestNilHead(t *testing.T) {
	if Nil.Head() != Nil {
		t.Fatalf(`expected Nil.Head() == Nil, got %s`, Nil.Head())
	}
}

func TestNilTail(t *testing.T) {
	if Nil.Tail() != Nil {
		t.Fatalf(`expected Nil.Tail() == Nil, got %s`, Nil.Tail())
	}
}

func TestNilPut(t *testing.T) {
	v, err := Nil.Put(Int(42))
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v.Head() != Int(42) {
		t.Fatalf(`expected v.Head() == Int(42), got %s`, v.Head())
	}
}

func TestNilEmpty(t *testing.T) {
	if Nil.Empty() == false {
		t.Fatalf(`expected Nil.Empty() == true, got false`)
	}
}

func TestNilString(t *testing.T) {
	if Nil.String() != "nil" {
		t.Fatalf(`expected Nil.String() == "nil", got %s`, Nil.String())
	}
}
