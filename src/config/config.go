package config

import (
	"encoding/json"
	"os"
	"time"
)

const version = "0.02.00"

type Config struct {
	StartTime time.Time
	Debug     bool
	token     string
	Version   string

	blacklistMap map[string]bool

	json struct {
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
			CmdBlacklist []string `json:"commandBlacklist"`
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
			CmdBlacklist []string `json:"commandBlacklist"`
		} `json:"debug"`
	}
}

func NewConfig(path string, debug bool) (*Config, error) {
	var cfg Config
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&(cfg.json))

	cfg.Debug = debug
	cfg.StartTime = time.Now()
	cfg.Version = version

	return &cfg, nil
}
