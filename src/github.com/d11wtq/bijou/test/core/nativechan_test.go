package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"sync"
	"testing"
)

func TestGoChanPortType(t *testing.T) {
	port := core.GoChanPort(make(core.ValueChannel), make(core.ValueChannel))

	if port.Type() != runtime.PortType {
		t.Fatalf(`expected port.Type() == PortType, got %s`, port.Type())
	}
}

func TestGoChanPortWriteAcceptAribtraryValues(t *testing.T) {
	var (
		res runtime.Value
		err error
		wg  sync.WaitGroup
	)

	for _, v := range []runtime.Value{
		runtime.True,
		runtime.Int(42),
		runtime.EmptyList,
		runtime.Symbol("foo"),
	} {
		this, that := core.GoChanPortPair()

		wg.Add(1)
		go func() {
			res, err = that.Accept()
			wg.Done()
		}()
		this.Write(v)
		wg.Wait()

		if err != nil {
			t.Fatalf(`expected err == nil, got %s`, err)
		}

		if !runtime.Eq(res, v) {
			t.Fatalf(`expected res == %s, got %s`, v, res)
		}
	}
}

func TestGoChanPortWritesIgnoredWhenClosed(t *testing.T) {
	var wg sync.WaitGroup

	this, that := core.GoChanPortPair()

	wg.Add(1)
	go func() {
		this.Close()
		that.Close()
		wg.Done()
	}()
	wg.Wait()
	err := this.Write(runtime.True)

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
}

func TestGoChanPortReturnsNilOnReadsWhenClosed(t *testing.T) {
	var wg sync.WaitGroup

	this, that := core.GoChanPortPair()

	wg.Add(1)
	go func() {
		this.Close()
		wg.Done()
	}()
	wg.Wait()
	res, err := that.Accept()

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if res != runtime.Nil {
		t.Fatalf(`expected res == Nil, got %s`, res)
	}
}
