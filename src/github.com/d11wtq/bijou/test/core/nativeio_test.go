package core_test

import (
	"bytes"
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestGoIoPortType(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	port := core.GoIoPort(nil, buf)

	if port.Type() != runtime.PortType {
		t.Fatalf(`expected port.Type() == PortType, got %s`, port.Type())
	}
}

func TestGoIoPortWritesStrings(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	port := core.GoIoPort(nil, buf)

	res, err := port.Write(runtime.String("example"))

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if res != port {
		t.Fatalf(`expected res == port, got %s`, res)
	}

	if buf.String() != "example" {
		t.Fatalf(`expected buf.String() == "example", got %s`, buf.String())
	}
}

func TestGoIoPortWritesNonStrings(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	port := core.GoIoPort(nil, buf)

	v := runtime.EmptyList.Append(runtime.Int(42))
	res, err := port.Write(v)

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if res != port {
		t.Fatalf(`expected res == port, got %s`, res)
	}

	if buf.String() != v.String() {
		t.Fatalf(`expected buf.String() == v.String(), got %s`, buf.String())
	}
}

func TestGoIoPortFailsWritesOnNilWriter(t *testing.T) {
	port := core.GoIoPort(nil, nil)
	res, err := port.Write(runtime.String("x"))

	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if res != nil {
		t.Fatalf(`expected res == nil, got %s`, res)
	}
}
