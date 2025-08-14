package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world"
	if result := Hello(); result != expected {
		t.Errorf("Ожидалось %q, получено %q", expected, result)
	}
}
