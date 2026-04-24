package game

import "testing"

func TestLevelsSetLevels(t *testing.T) {
	l := NewLevels()
	if err := l.SetLevels(); err != nil {
		t.Fatalf("expected no error from SetLevels, got %v", err)
	}

	if len(l.layers) == 0 {
		t.Fatal("expected level layers to be present")
	}

	if len(l.layers) != 2 {
		t.Fatalf("expected 2 layers, got %d", len(l.layers))
	}

	if len(l.layers[0]) == 0 {
		t.Fatal("expected first layer to have tiles")
	}
}
