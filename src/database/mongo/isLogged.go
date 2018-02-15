package mongo

import "gopkg.in/mgo.v2/bson"

func (db *Mongo) IsLogged(mid string) bool {
	ct, err := db.messages.Find(bson.M{"id": mid}).Count()
	return err == nil && ct > 0
}
