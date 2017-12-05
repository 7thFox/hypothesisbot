package mongo

import (
	"github.com/bwmarrin/discordgo"
)

func (db *Mongo) LogMessage(m *discordgo.Message) error {
	c := db.session.DB("hypothesis_bot").C("messages")
	err := c.Insert(m)
	return err
}
