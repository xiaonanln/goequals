package goequals

import (
	"testing"
)

func TestEquals(t *testing.T) {
	mustEqual(t, 1, 1)
	mustEqual(t, int8(1), int64(1))
	mustEqual(t, uint64(1), int64(1))
	mustNotEqual(t, 1, 2)
	mustNotEqual(t, uint64(0xffffffffffffffff), int64(-1))
	mustEqual(t, true, true)
	mustEqual(t, false, false)
	mustNotEqual(t, false, true)
	mustNotEqual(t, true, false)

	mustEqual(t, []int{1, 2, 3}, []int{1, 2, 3})
	mustEqual(t, []int{1, 2, 3}, []int8{1, 2, 3})
	mustEqual(t, []int{1, 2, 3}, []uint64{1, 2, 3})
	mustNotEqual(t, []int{1, 2, 3}, []int{1, 2, 4})
	mustNotEqual(t, []int{1, 2, 3}, []int{1, 2, 3, 4})

	mustEqual(t, map[string]int{}, map[string]int{})
	mustEqual(t, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2})
	mustEqual(t, map[interface{}]uint{"a": 1, "b": 2}, map[interface{}]int{"a": 1, "b": 2})

	mustEqual(t, [1]int{1}, [1]int{1})
	mustNotEqual(t, [1]int{1}, [2]int{1, 2})

	mustEqual(t, complex(1, 1), complex(1, 1))
	mustNotEqual(t, complex(1, 1), complex(1, 2))
}

func mustEqual(t *testing.T, v1, v2 interface{}) {
	if Equals(v1, v2) {

	} else {
		t.Fatalf("%T(%v) should equal to %T(%v), but Equals returns false", v1, v1, v2, v2)
	}
}

func mustNotEqual(t *testing.T, v1, v2 interface{}) {
	if Equals(v1, v2) {
		t.Fatalf("%T(%v) should not equal to %T(%v), but Equals returns true", v1, v1, v2, v2)
	} else {

	}
}
