package acceptance_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestWrite(t *testing.T) {
	buf := test.FakeIO(nil)
	port := core.GoIoPort(nil, buf)
	env := core.RootEnv()
	env.Def("test-port", port)

	AssertRunEqualWithEnv(t, `(-> test-port "foo")`, port, env)

	if buf.String() != "foo" {
		t.Fatalf(`expected buf == "foo", got %s`, buf.String())
	}
}
