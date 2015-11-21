package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestGoIoPortType(t *testing.T) {
	buf := test.FakeIO(nil)
	port := core.GoIoPort(nil, buf)

	if port.Type() != runtime.PortType {
		t.Fatalf(`expected port.Type() == PortType, got %s`, port.Type())
	}
}

func TestGoIoPortWritesStrings(t *testing.T) {
	buf := test.FakeIO(nil)
	port := core.GoIoPort(nil, buf)

	err := port.Write(runtime.String("example"))

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if buf.String() != "example" {
		t.Fatalf(`expected buf.String() == "example", got %s`, buf.String())
	}
}

func TestGoIoPortWritesNonStrings(t *testing.T) {
	buf := test.FakeIO(nil)
	port := core.GoIoPort(nil, buf)

	v := runtime.EmptyList.Append(runtime.Int(42))
	err := port.Write(v)

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if buf.String() != v.String() {
		t.Fatalf(`expected buf.String() == v.String(), got %s`, buf.String())
	}
}

func TestGoIoPortFailsWritesOnNilWriter(t *testing.T) {
	port := core.GoIoPort(nil, nil)
	err := port.Write(runtime.String("x"))

	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}

func TestGoIoPortAcceptsInts(t *testing.T) {
	buf := test.FakeIO([]byte("foo"))
	port := core.GoIoPort(buf, nil)

	for _, wanted := range []runtime.Value{
		runtime.Int(102),
		runtime.Int(111),
		runtime.Int(111),
		runtime.Nil,
	} {
		actual, err := port.Accept()
		if err != nil {
			t.Fatalf(`expected err == nil, got %s`, err)
		}
		if actual != wanted {
			t.Fatalf(`expected v == %s, got %s`, wanted, actual)
		}
	}
}

func TestGoIoPortFailsAcceptOnNilReader(t *testing.T) {
	port := core.GoIoPort(nil, nil)
	_, err := port.Accept()

	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}

func TestGoIoPortClosesBothReaderAndWriter(t *testing.T) {
	bufr := test.FakeIO(nil)
	bufw := test.FakeIO(nil)
	port := core.GoIoPort(bufr, bufw)

	err := port.Close()

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if bufr.Closed == false {
		t.Fatalf(`expected bufr.Closed, got false`)
	}
	if bufw.Closed == false {
		t.Fatalf(`expected bufw.Closed, got false`)
	}
}
