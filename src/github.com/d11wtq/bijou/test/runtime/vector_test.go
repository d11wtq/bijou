package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func AssertVectorContains(t *testing.T, vec *Vector, elems map[uint32]Value) {
	for k, v := range elems {
		x, ok := vec.Find(k)
		if ok == false {
			t.Fatalf(`expected vec.Find(%d) to be ok, got false`, k)
		}
		if !Eq(x, v) {
			t.Fatalf(`expected vec.Find(%d) == %s, got %s`, k, v, x)
		}
	}
}

func TestVectorFind1Deep(t *testing.T) {
	vec := &Vector{
		Elements: []Value{Int(42), Int(21), Int(17)},
		Shift:    0, // 5 * (1 - 1)
	}

	x, ok := vec.Find(0)
	if ok == false {
		t.Fatalf(`expected vec.Find(0) to be ok, got false`)
	}
	if x != Int(42) {
		t.Fatalf(`expected vec.Find(0) == Int(42), got %s`, x)
	}

	y, ok := vec.Find(1)
	if ok == false {
		t.Fatalf(`expected vec.Find(1) to be ok, got false`)
	}
	if y != Int(21) {
		t.Fatalf(`expected vec.Find(1) == Int(21), got %s`, y)
	}

	z, ok := vec.Find(2)
	if ok == false {
		t.Fatalf(`expected vec.Find(2) to be ok, got false`)
	}
	if z != Int(17) {
		t.Fatalf(`expected vec.Find(2) == Int(17), got %s`, z)
	}

	a, ok := vec.Find(3)
	if ok == true {
		t.Fatalf(`expected vec.Find(3) not to be ok, got true`)
	}
	if a != nil {
		t.Fatalf(`expected vec.Find(3) == nil, got %s`, a)
	}
}

func TestVectorFind2Deep(t *testing.T) {
	vec := &Vector{
		Elements: []Value{
			&Vector{Elements: []Value{Int(42), Int(21), Int(17)}},
		},
		Shift: 5, // 5 * (2 - 1)
	}

	x, ok := vec.Find(0)
	if ok == false {
		t.Fatalf(`expected vec.Find(0) to be ok, got false`)
	}
	if x != Int(42) {
		t.Fatalf(`expected vec.Find(0) == Int(42), got %s`, x)
	}

	y, ok := vec.Find(1)
	if ok == false {
		t.Fatalf(`expected vec.Find(1) to be ok, got false`)
	}
	if y != Int(21) {
		t.Fatalf(`expected vec.Find(1) == Int(21), got %s`, y)
	}

	z, ok := vec.Find(2)
	if ok == false {
		t.Fatalf(`expected vec.Find(2) to be ok, got false`)
	}
	if z != Int(17) {
		t.Fatalf(`expected vec.Find(2) == Int(17), got %s`, z)
	}

	a, ok := vec.Find(3)
	if ok == true {
		t.Fatalf(`expected vec.Find(3) not to be ok, got true`)
	}
	if a != nil {
		t.Fatalf(`expected vec.Find(3) == nil, got %s`, a)
	}
}

func TestVectorUpdateViaSet1Deep(t *testing.T) {
	vec := &Vector{
		Elements: []Value{Int(42), Int(21), Int(17)},
		Shift:    0, // 5 * (1 - 1)
	}

	cpy, ok := vec.Set(1, Int(57))
	if ok == false {
		t.Fatalf(`expected vec.Set(1, ...) to be ok, got false`)
	}

	AssertVectorContains(
		t, cpy,
		map[uint32]Value{
			0: Int(42),
			1: Int(57),
			2: Int(17),
		},
	)

	AssertVectorContains(
		t, vec,
		map[uint32]Value{
			0: Int(42),
			1: Int(21),
			2: Int(17),
		},
	)

	_, ok = cpy.Find(3)
	if ok == true {
		t.Fatalf(`expected cpy.Find(3) not to be ok, got true`)
	}
}

func TestVectorUpdateViaSet2Deep(t *testing.T) {
	vec := &Vector{
		Elements: []Value{
			&Vector{Elements: []Value{Int(42), Int(21), Int(17)}},
		},
		Shift: 5, // 5 * (2 - 1)
	}

	cpy, ok := vec.Set(1, Int(57))
	if ok == false {
		t.Fatalf(`expected vec.Set(1, ...) to be ok, got false`)
	}

	AssertVectorContains(
		t, cpy,
		map[uint32]Value{
			0: Int(42),
			1: Int(57),
			2: Int(17),
		},
	)

	AssertVectorContains(
		t, vec,
		map[uint32]Value{
			0: Int(42),
			1: Int(21),
			2: Int(17),
		},
	)

	_, ok = cpy.Find(3)
	if ok == true {
		t.Fatalf(`expected cpy.Find(3) not to be ok, got true`)
	}
}

func TestVectorAppendViaSet1Deep(t *testing.T) {
	vec := &Vector{
		Elements: []Value{Int(42), Int(21), Int(17)},
		Shift:    0, // 5 * (1 - 1)
	}

	cpy, ok := vec.Set(3, Int(57))
	if ok == false {
		t.Fatalf(`expected vec.Set(3, ...) to be ok, got false`)
	}

	AssertVectorContains(
		t, cpy,
		map[uint32]Value{
			0: Int(42),
			1: Int(21),
			2: Int(17),
			3: Int(57),
		},
	)

	_, ok = cpy.Find(4)
	if ok == true {
		t.Fatalf(`expected cpy.Find(4) not to be ok, got true`)
	}

	_, ok = vec.Find(3)
	if ok == true {
		t.Fatalf(`expected vec.Find(3) not to be ok, got true`)
	}
}

func TestVectorAppendViaSet2Deep(t *testing.T) {
	vec := &Vector{
		Elements: []Value{
			&Vector{Elements: []Value{Int(42), Int(21), Int(17)}},
		},
		Shift: 5, // 5 * (2 - 1)
	}

	cpy, ok := vec.Set(3, Int(57))
	if ok == false {
		t.Fatalf(`expected vec.Set(3, ...) to be ok, got false`)
	}

	AssertVectorContains(
		t, cpy,
		map[uint32]Value{
			0: Int(42),
			1: Int(21),
			2: Int(17),
			3: Int(57),
		},
	)

	_, ok = cpy.Find(4)
	if ok == true {
		t.Fatalf(`expected cpy.Find(4) not to be ok, got true`)
	}

	_, ok = vec.Find(3)
	if ok == true {
		t.Fatalf(`expected vec.Find(3) not to be ok, got true`)
	}
}

func TestVectorAppendOverflow1Deep(t *testing.T) {
	elems := make([]Value, 0, 32)
	for i := 0; i < 32; i += 1 {
		elems = append(elems, Int(i))
	}
	vec := &Vector{Elements: elems, Shift: 0} // 5 * (1 - 1)

	cpy, ok := vec.Set(32, Int(32))
	if ok == false {
		t.Fatalf(`expected vec.Set(32, ...) to be ok, got false`)
	}

	AssertVectorContains(
		t, cpy,
		map[uint32]Value{
			0:  Int(0),
			31: Int(31),
			32: Int(32),
		},
	)

	_, ok = cpy.Find(33)
	if ok == true {
		t.Fatalf(`expected cpy.Find(33) not to be ok, got true`)
	}
}

func TestVectorAppendOverflow2Deep(t *testing.T) {
	nodes := make([]Value, 0, 32)
	for i := 0; i < 32; i += 1 {
		elems := make([]Value, 0, 32)
		for j := 0; j < 32; j += 1 {
			elems = append(elems, Int(i*32+j))
		}
		nodes = append(nodes, &Vector{Elements: elems, Shift: 0})
	}

	vec := &Vector{Elements: nodes, Shift: 5} // 5 * (2 - 1)

	cpy, ok := vec.Set(1024, Int(1024))
	if ok == false {
		t.Fatalf(`expected vec.Set(1024, ...) to be ok, got false`)
	}

	AssertVectorContains(
		t, cpy,
		map[uint32]Value{
			0:    Int(0),
			31:   Int(31),
			32:   Int(32),
			1023: Int(1023),
			1024: Int(1024),
		},
	)

	_, ok = cpy.Find(1025)
	if ok == true {
		t.Fatalf(`expected cpy.Find(1025) not to be ok, got true`)
	}
}
