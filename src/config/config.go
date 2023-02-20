package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/amarwaits/Syndica-THT/src/model"
)

type Config struct {
	ServerAddress       string `json:"serverAddress"`
	InitialDataFilePath string `json:"initialDataFilePath"`
}

func NewConfig() *Config {
	return &Config{
		ServerAddress:       ":8080",
		InitialDataFilePath: "assets/movies.json",
	}
}

func (c *Config) LoadMoviesFromFile() ([]model.Movie, error) {
	data, err := ioutil.ReadFile(c.InitialDataFilePath)
	if err != nil {
		return nil, err
	}

	var movies []model.Movie
	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
