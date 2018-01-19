package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/7thFox/hypothesisbot/command"
	"github.com/7thFox/hypothesisbot/database"
	"github.com/7thFox/hypothesisbot/log"
)

const version = "0.01.00"

type Config struct {
	json configJSON

	db database.Database

	cmds map[string]command.Command

	lgr           *log.MultiLogger
	lgrHasSession bool

	StartTime time.Time
	Debug     bool
	token     string
	Version   string
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
		Logging struct {
			Console   bool   `json:"console"`
			ChannelID string `json:"channel"`
			Dbname    string `json:"db"`
		} `json:"log"`
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
		Logging struct {
			Console   bool   `json:"console"`
			ChannelID string `json:"channel"`
			Dbname    string `json:"db"`
		} `json:"log"`
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
	cfg.StartTime = time.Now()
	cfg.Version = version

	return &cfg
}
