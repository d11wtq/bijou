package acceptance_test

import (
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestSpawn(t *testing.T) {
	AssertRunType(
		t,
		`
		(def test
		  (fn (p)))
		(spawn test)
		`,
		runtime.PortType,
	)

	AssertRunEqual(
		t,
		`
		(def ping-pong
		  (fn (p)
		    (if (= (receive! p) 'ping)
			  (send! p 'pong)
			  (send! p 'ping))
		    (ping-pong p)))

		(def p (spawn ping-pong))

		(def message!
		  (fn (p msg)
		    (send! p msg)
			(receive! p)))

		(def ping
		  (fn (p) (message! p 'ping)))

		(def pong
		  (fn (p) (message! p 'pong)))

		(list (ping p) (pong p))
		`,
		runtime.EmptyList.
			Append(runtime.Symbol("pong")).
			Append(runtime.Symbol("ping")),
	)
}
