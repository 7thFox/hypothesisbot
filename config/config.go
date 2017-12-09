package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/7thFox/hypothesisbot/database"
)

type Config struct {
	json configJSON

	db    database.Database
	Debug bool
	token string
}

type configJSON struct {
	Global struct {
		TokenPath  string   `json:"token"`
		Prefix     string   `json:"prefix"`
		LogServers []string `json:"logservers"`
		Database   struct {
			Dbtype string `json:"type"`
			Host   string `json:"host"`
			Dbname string `json:"dbname"`
		} `json:"db"`
	} `json:"global"`
	Debug struct {
		TokenPath  string   `json:"token"`
		Prefix     string   `json:"prefix"`
		LogServers []string `json:"logservers"`
		Database   struct {
			CopyProduction bool   `json:"copyprod"`
			Dbtype         string `json:"type"`
			Host           string `json:"host"`
			Dbname         string `json:"dbname"`
		} `json:"db"`
	} `json:"debug"`
}

func NewConfig(path string, d bool) *Config {
	var cfg Config
	fmt.Printf("Loading Config from %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&(cfg.json))

	cfg.Debug = d

	return &cfg
}
