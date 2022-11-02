package rtree

import "testing"

func TestOneD_Insert(t *testing.T) {
	rt := NewOneD[int, string]()

	x := 1
	y := 10
	z := 5
	value := "hello"
	replacedValue := "world"

	rt.Insert(x, y, value)
	if rt.Len() != 1 {
		t.Error("IntRTree length unexpected after insert, expected 1")
	}

	rt.Search(z, z, func(min, max int, data string) bool {
		if data != value {
			t.Error("IntRTree search returned incorrect value")
		}
		return true
	})

	rt.Replace(x, y, value, x, y, replacedValue)
	if rt.Len() != 1 {
		t.Error("IntRTree length unexpected after replace, expected 1")
	}

	rt.Search(x, y, func(min, max int, data string) bool {
		if data != replacedValue {
			t.Error("IntRTree search returned incorrect value")
		}
		return true
	})

	rt.Delete(x, y, replacedValue)
	if rt.Len() != 0 {
		t.Error("IntRTree length unexpected after delete, expected 0")
	}

	rt.Insert(x, y, value)
	rt.Insert(x, y, replacedValue)
	if rt.Len() != 2 {
		t.Error("IntRTree length unexpected after replace, expected 2")
	}

	rt.Clear()
	if rt.Len() != 0 {
		t.Error("IntRTree length unexpected after clear, expected 0")
	}
}
