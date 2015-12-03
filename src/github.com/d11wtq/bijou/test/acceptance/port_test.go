package acceptance_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestSend(t *testing.T) {
	buf := test.FakeIO(nil)
	port := core.GoIoPort(nil, buf)
	env := core.RootEnv()
	env.Def("test-port", port)

	AssertRunEqualWithEnv(t, `(send! test-port "foo")`, port, env)

	if buf.String() != "foo" {
		t.Fatalf(`expected buf == "foo", got %s`, buf.String())
	}
}

func TestReceive(t *testing.T) {
	buf := test.FakeIO([]byte("foo"))
	port := core.GoIoPort(buf, nil)
	env := core.RootEnv()
	env.Def("test-port", port)

	AssertRunEqualWithEnv(t, `(receive! test-port)`, runtime.Int(102), env)

	if buf.String() != "oo" {
		t.Fatalf(`expected buf == "oo", got %s`, buf.String())
	}
}
