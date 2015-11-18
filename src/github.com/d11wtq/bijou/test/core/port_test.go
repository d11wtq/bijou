package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestWriteToPort(t *testing.T) {
	buf := test.FakeIO(nil)
	port := core.GoIoPort(nil, buf)

	args := runtime.EmptyList.
		Append(port).
		Append(runtime.String("example"))

	v, err := core.Write(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if !v.Eq(port) {
		t.Fatalf(`expected v == port, got %s`, v)
	}
	if buf.String() != "example" {
		t.Fatalf(`expected buf.String() == "example", got %s`, buf.String())
	}
}
