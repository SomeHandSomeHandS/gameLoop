package models

type Config struct {
	TestIt       bool   `json:"testit"`
	URI          string `json:"uri"`
	ScreenWidth  int    `json:"screenWidth"`
	ScreenHeight int    `json:"screenHeight"`
	TileSize     int    `json:"tileSize"`
}

type LevelBook struct {
	Levels []Level
}

type Level struct {
	Name   string  `json:"name"`
	Layers [][]int `json:"layers"`
}
