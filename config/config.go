package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	json configJSON

	Debug bool
	token string
}

type configJSON struct {
	Global struct {
		TokenPath string `json:"token"`
		Database  struct {
			Dbtype   string `json:"type"`
			Location string `json:"location"`
		}
	} `json:"global"`
	Debug struct {
		TokenPath string `json:"token"`
		Database  struct {
			CopyProduction bool   `json:"copyprod"`
			Dbtype         string `json:"type"`
			Location       string `json:"location"`
		}
	} `json:"debug"`
}

func NewConfig(path string, d bool) *Config {
	var cfg Config
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&(cfg.json))

	cfg.Debug = d
	return &cfg
}
