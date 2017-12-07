package mongo

import (
	"github.com/bwmarrin/discordgo"
)

func (db *Mongo) LogMessage(m *discordgo.Message) error {
	err := db.messages.Insert(m)
	return err
}
