package game

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const(
	spriteSize = 10
)

type Sprites struct {
	spriteImage *ebiten.Image
}

func InitSprites() (*Sprites, error) {
	s := &Sprites{}
	err := s.init("")
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Sprites) init(spritePath string) error {
	// TODO wtf is this? I don't think we need to load a sprite sheet for this game, but if we do, we can load it here.
	if spritePath == "" {
		spritePath = "adventurer_sprite_sheet_v1.1.png"
	}

	f, err := os.Open(spritePath)
	if err != nil {
		log.Fatalf("failed to open sprite image %q: %v", spritePath, err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("failed to decode sprite image %q: %v", spritePath, err)
	}
	s.spriteImage = ebiten.NewImageFromImage(img)
	return nil
}

func (s *Sprites) GetSpriteImage() *ebiten.Image {

	return s.spriteImage
}