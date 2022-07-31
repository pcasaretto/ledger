package monetary

import (
	"testing"
)

func TestBasics(t *testing.T) {
	a := NewInt(1)
	b := NewInt(1)
	c := NewInt(2)
	d := NewIntFromString("10000000000000000000000000000000000000000000000000000001")
	e := NewIntFromString("10000000000000000000000000000000000000000000000000000002")

	if !a.Eq(a) {
		t.Errorf("a.Eq(a) = false")
	}

	if !a.Eq(b) {
		t.Errorf("a.Eq(b) = false")
	}

	if a.Eq(c) {
		t.Errorf("a.Eq(c) = true")
	}

	if a.Gte(c) {
		t.Errorf("a.Gte(c) = false")
	}

	if !a.Lt(c) {
		t.Errorf("a.Lt(c) = false")
	}

	if !a.Lte(c) {
		t.Errorf("a.Lte(c) = false")
	}

	if !d.Add(a).Eq(e) {
		t.Errorf("d.Add(a).Eq(e) = false")
	}

	if !e.Sub(d).Eq(a) {
		t.Errorf("e.Sub(d).Eq(a) = false")
	}
}
