package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type NewMessage struct {
	ID        string `json:"id"`
	ChannelID string `json:"channelid"`
}

// func (db *Mongo) NewestMessages() (map[string]string, error) {

// 	itr := db.messages.Pipe([]bson.M{
// 		{"$group": bson.M{
// 			"_id":       "$channelid",
// 			"channelid": bson.M{"$first": "$channelid"},
// 			"id":        bson.M{"$max": "$id"},
// 		}},
// 	}).Iter()

// 	m := NewMessage{}
// 	ms := map[string]string{}
// 	for itr.Next(&m) {
// 		ms[m.ChannelID] = m.ID
// 	}

// 	return ms, nil
// }

func (db *Mongo) NewestMessagesBefore(t time.Time) (map[string]string, error) {

	itr := db.messages.Pipe([]bson.M{
		{"$match": bson.M{"timestamp": bson.M{"$lt": t}}},
		{"$group": bson.M{
			"_id":       "$channelid",
			"channelid": bson.M{"$first": "$channelid"},
			"id":        bson.M{"$max": "$id"},
		}},
	}).Iter()

	m := NewMessage{}
	ms := map[string]string{}
	for itr.Next(&m) {
		ms[m.ChannelID] = m.ID
	}

	return ms, nil
}
