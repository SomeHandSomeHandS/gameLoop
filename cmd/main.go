package main

import (
	"encoding/json"
	"os"

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

func loadConfig(filename string) (config models.Config, err error) {
	config = models.Config{}

	// Read the file
	data, err := os.ReadFile(filename)
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

	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	// load config
	_, err = loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Cool game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
