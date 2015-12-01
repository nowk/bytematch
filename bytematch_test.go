package bytematch

import (
	"testing"
)

func TestExactMatch(t *testing.T) {
	var (
		a = []byte("Hello World!")
		b = []byte("llo")
	)

	i, m := Compare(a, b)
	if exp := 2; exp != i {
		t.Errorf("expected %d, got %d", exp, i)
	}
	if exp := Exact; exp != m {
		t.Errorf("expected %d, got %d", exp, m)
	}
}

func TestPartialMatch(t *testing.T) {
	var (
		a = []byte("Hello Wor")
		b = []byte("World!")
	)

	i, m := Compare(a, b)
	if exp := 6; exp != i {
		t.Errorf("expected %d, got %d", exp, i)
	}
	if exp := Partial; exp != m {
		t.Errorf("expected %d, got %d", exp, m)
	}
}

func TestNotMatch(t *testing.T) {
	var (
		a = []byte("Hello World!")
		b = []byte("foo")
	)

	i, m := Compare(a, b)
	if exp := -1; exp != i {
		t.Errorf("expected %d, got %d", exp, i)
	}
	if exp := Not; exp != m {
		t.Errorf("expected %d, got %d", exp, m)
	}
}
