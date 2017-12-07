package mongo

import "gopkg.in/mgo.v2/bson"

func (db *Mongo) IsLogged(mid string) bool {
	c := db.db.C("messages")
	ct, err := c.Find(bson.M{"id": mid}).Count()
	return err == nil && ct > 0
}
