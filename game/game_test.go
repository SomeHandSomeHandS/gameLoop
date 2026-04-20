package game

import "testing"

func TestLayout(t *testing.T) {
    g := &Game{}
    width, height := g.Layout(999, 999)

    if width != screenWidth || height != screenHeight {
        t.Fatalf("expected layout %dx%d, got %dx%d", screenWidth, screenHeight, width, height)
    }
}
