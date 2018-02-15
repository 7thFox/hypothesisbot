package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

func (m *Mongo) MsgCount(by string, aggregate bool) ([]bson.M, error) {
	var id bson.M
	var sort bson.M
	if aggregate {
		id = bson.M{
			"$" + by: bson.M{"$dateFromString": bson.M{"dateString": "$timestamp", "timezone": "+00"}},
		}
		sort = bson.M{"_id": 1}
	} else {
		id = bson.M{
			"year": bson.M{"$year": bson.M{"$dateFromString": bson.M{"dateString": "$timestamp", "timezone": "+00"}}},
			by:     bson.M{"$" + by: bson.M{"$dateFromString": bson.M{"dateString": "$timestamp", "timezone": "+00"}}},
		}
		sort = bson.M{"_id.year": 1, "_id." + by: 1}
	}

	pipe := m.messages.Pipe([]bson.M{
		{"$group": bson.M{
			"_id":   id,
			"count": bson.M{"$sum": 1},
		}},
		{"$sort": sort},
	})
	resp := []bson.M{}
	err := pipe.All(&resp)
	return resp, err
}
