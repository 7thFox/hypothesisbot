package database

import (
	"errors"
	"time"

	"github.com/7thFox/hypothesisbot/database/mongo"
	"github.com/bwmarrin/discordgo"
)

type Database interface {
	LogMessage(m *discordgo.Message) error
	IsLogged(mid string) bool
	OldestMessageInChannel(cid string) (*discordgo.Message, error)
	NewestMessageInChannel(cid string) (*discordgo.Message, error)
	NewestMessagesBefore(time.Time) (map[string]string, error)
}

func NewDatabase(t string, host string, name string) (interface{}, error) {
	if t == "mongo" {
		return mongo.NewMongo(host, name)
	}
	return nil, errors.New("Unsupported Database type")
}
