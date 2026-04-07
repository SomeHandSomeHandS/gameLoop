package models

type Config struct {
	TestIt        bool   `json:"testit"`
	URI           string `json:"uri"`
	MongoOpenTime int    `json:"mongoOpenTime"`
}
