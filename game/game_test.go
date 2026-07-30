package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestLayout(t *testing.T) {
	g := &Game{}
	width, height := g.Layout(999, 999)

	if width != screenWidth || height != screenHeight {
		t.Fatalf("expected layout %dx%d, got %dx%d", screenWidth, screenHeight, width, height)
	}
}

func TestShouldEndGame(t *testing.T) {
	if !shouldEndGame(func(key ebiten.Key) bool {
		return key == ebiten.KeyE
	}) {
		t.Fatal("expected E to end the game")
	}

	if shouldEndGame(func(key ebiten.Key) bool {
		return key == ebiten.KeySpace
	}) {
		t.Fatal("space should not end the game")
	}
}
