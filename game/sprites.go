package game

type Sprites struct {
}

func InitSprites() *Sprites {
	return &Sprites{}
}

// func init() {
// 	imgFile := "adventurer_sprite_sheet_v1.1.png"
// 	f, err := os.Open(imgFile)
// 	if err != nil {
// 		log.Fatalf("failed to open sprite image %q: %v", imgFile, err)
// 	}
// 	defer f.Close()

// 	img, _, err := image.Decode(f)
// 	if err != nil {
// 		log.Fatalf("failed to decode sprite image %q: %v", imgFile, err)
// 	}
// 	tilesImage = ebiten.NewImageFromImage(img)
// }
