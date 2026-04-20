package game

import (
	"reflect"
	"testing"
)

func TestLevelsSetLevels(t *testing.T) {
	l := &Levels{}
	levelBook := l.SetLevels()

	if levelBook == nil {
		t.Fatal("expected non-nil LevelBook")
	}

	if len(levelBook.Levels) != 1 {
		t.Fatalf("expected 1 level, got %d", len(levelBook.Levels))
	}

	level := levelBook.Levels[0]
	if level.Name != "COOL TEST" {
		t.Fatalf("expected level name %q, got %q", "COOL TEST", level.Name)
	}

	gotLayers := GetLevel(level)
	if !reflect.DeepEqual(gotLayers, level.Layers) {
		t.Fatal("GetLevel returned unexpected layer data")
	}

	if len(gotLayers) == 0 {
		t.Fatal("expected level layers to be present")
	}
}
