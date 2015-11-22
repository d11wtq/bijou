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
		    (if (= (<- p) 'ping)
			  (-> p 'pong)
			  (-> p 'ping))
		    (ping-pong p)))

		(def p (spawn ping-pong))

		(def ping
		  (fn (p)
		    (-> p 'ping)
			(<- p)))

		(def pong
		  (fn (p)
		    (-> p 'pong)
			(<- p)))

		(list (ping p) (pong p))
		`,
		runtime.EmptyList.
			Append(runtime.Symbol("pong")).
			Append(runtime.Symbol("ping")),
	)
}
