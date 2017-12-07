package mongo

import (
	"github.com/bwmarrin/discordgo"
	"gopkg.in/mgo.v2/bson"
)

func (db *Mongo) OldestMessageInChannel(cid string) (*discordgo.Message, error) {
	m := discordgo.Message{}
	err := db.messages.Find(bson.M{"channelid": cid}).Sort("timestamp").Limit(1).One(&m)
	return &m, err
}

func (db *Mongo) NewestMessageInChannel(cid string) (*discordgo.Message, error) {
	m := discordgo.Message{}
	err := db.messages.Find(bson.M{"channelid": cid}).Sort("timestamp-").Limit(1).One(&m)
	return &m, err
}
