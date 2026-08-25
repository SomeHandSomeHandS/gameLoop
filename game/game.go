package game

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mode int

// Game represents the game state and logic.
type Game struct {
	board   *Board
	input   *Input
	levels  *Levels
	sprites *Sprites

	mode          Mode
	gameOverCount int

	touchIDs   []ebiten.TouchID
	gamepadIDs []ebiten.GamepadID
	keys       []ebiten.Key
}

var (
	tilesImage, spriteImage *ebiten.Image
)

const (
	screenWidth  = 240
	screenHeight = 240
	boardSize    = 4
)

const (
	tileSize        = 16
	ModeTitle  Mode = iota
	ModeGame
	ModeGameOver
)

func init() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)

}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {

	var err error

	g := &Game{
		mode: ModeTitle,
	}

	g.input, err = NewInput()
	if err != nil {
		return nil, err
	}

	g.levels, err = InitLevels()
	if err != nil {
		return nil, err
	}
	err = g.levels.getLevels()
	if err != nil {
		return nil, err
	}

	g.board, err = NewBoard(boardSize)
	if err != nil {
		return nil, err
	}

	g.sprites, err = InitSprites()
	if err != nil {
		return nil, err
	}

	spriteImage = g.sprites.spriteImage

	return g, nil
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		if g.isKeyJustPressed() {
			g.mode = ModeGame
		}

	case ModeGame:
		if shouldEndGame(inpututil.IsKeyJustPressed) {
			g.mode = ModeGameOver
			return nil
		}

		if g.isKeyJustPressed() {
			// update game state and layers
			g.levels.setLevel()

		}
		// update game state

	case ModeGameOver:
		return ebiten.Termination
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / tileSize

	// Draw each tile with each DrawImage call.
	// As the source images of all DrawImage calls are always same,
	// this rendering is done very efficiently.
	// For more detail, see https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Image.DrawImage
	const xCount = screenWidth / tileSize
	// TODO wtf levels array
	for _, l := range g.levels.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
