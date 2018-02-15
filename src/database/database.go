package database

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/mgo.v2/bson"
)

type Database interface {
	MsgCount(by string, aggregate bool) ([]bson.M, error)

	LogMessage(m *discordgo.Message) error
	IsLogged(mid string) bool
	OldestMessageInChannel(cid string) (*discordgo.Message, error)
	NewestMessageInChannel(cid string) (*discordgo.Message, error)
	NewestMessagesBefore(time.Time) (map[string]string, error)
}
