package database

import (
	"errors"

	"github.com/7thFox/hypothesisbot/database/sqlite"
	"github.com/bwmarrin/discordgo"
)

type Database interface {
	LogMessage(m *discordgo.MessageCreate)
}

func NewDatabase(t string, l string) (interface{}, error) {
	if t == "sqlite" {
		return sqlite.NewSqlite(l), nil
	}
	return nil, errors.New("Unsupported Database type")
}
