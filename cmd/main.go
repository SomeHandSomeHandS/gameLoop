package main

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"mongoConnector/game"
	"mongoConnector/models"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mode int

const (
	screenWidth  = 240
	screenHeight = 240
)

const (
	tileSize       = 16
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

var (
	tilesImage *ebiten.Image
)

func loadConfig(filename string) (config models.Config, err error) {
	config = models.Config{}

	// Read the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	// Unmarshal the JSON into a Config struct

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func main() {

	g := &game.Game{}
	l := &game.Levels{}

	// load config
	_, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// startup and get levels and set game
	levelBook := l.SetLevels()

	g.Layers = levelBook.Levels[0].Layers

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Cool game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
