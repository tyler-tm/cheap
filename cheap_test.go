package cheap

import (
	"sort"
	"testing"
)

func TestPushOnce(t *testing.T) {
	m := NewCHeap()

	want := "want"
	m.Push(want, 1)

	got, _ := m.Peek()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPushMore(t *testing.T) {
	m := NewCHeap()

	want := "yes"
	items := []struct {
		string
		int
	}{
		{"no", 1},
		{"no", 8},
		{"no", 3},
		{want, 13},
		{"no", 5},
		{"no", 2},
	}
	for _, v := range items {
		m.Push(v.string, v.int)
	}

	got, _ := m.Peek()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPopOnce(t *testing.T) {
	m := NewCHeap()
	want := "want"

	m.Push(want, 1)
	got, _ := m.Pop()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPopMore(t *testing.T) {
	m := NewCHeap()

	items := []struct {
		string
		int
	}{
		{"fourth", 1},
		{"second", 8},
		{"first", 13},
		{"third", 5},
	}
	for _, v := range items {
		m.Push(v.string, v.int)
	}
	sort.Slice(items, func(a int, b int) bool {
		return items[a].int > items[b].int
	})

	for _, want := range items {
		got, _ := m.Pop()
		if got != want.string {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestPopEmpty(t *testing.T) {
	m := NewCHeap()
	got, err := m.Pop()
	if got != nil {
		t.Errorf("Result returned when Popping from empty heap")
	}
	if err == nil {
		t.Errorf("No error thrown for Popping from empty heap")
	}
}

func TestPeekEmpty(t *testing.T) {
	m := NewCHeap()
	got, err := m.Peek()
	if got != nil {
		t.Errorf("Result returned when Peeking from empty heap")
	}
	if err == nil {
		t.Errorf("No error thrown for Peeking from empty heap")
	}
}
