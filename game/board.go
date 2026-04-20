package game

import "github.com/hajimehoshi/ebiten/v2"

type Board struct {
	size  int
	tiles map[*Tile]struct{}
}

// NewBoard generates a new Board with giving a size.
func NewBoard(size int) (*Board, error) {
	b := &Board{
		size:  size,
		tiles: map[*Tile]struct{}{},
	}

	return b, nil
}

// Update updates the board state.
func (b *Board) Update(input *Input) error {
	for t := range b.tiles {
		if err := t.Update(); err != nil {
			return err
		}
	}

	// if dir, ok := input.Dir(); ok {

	// }
	return nil
}

// Size returns the board size.
func (b *Board) Size() (int, int) {
	return b.size, b.size
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {

}
