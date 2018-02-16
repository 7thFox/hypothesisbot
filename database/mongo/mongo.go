package mongo

import (
	"gopkg.in/mgo.v2"
)

type Mongo struct {
	session  *mgo.Session
	messages *mgo.Collection
	db       *mgo.Database
}

func NewMongo(host string, name string) (*Mongo, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	s := Mongo{}
	s.session = session
	s.db = s.session.DB(name)
	s.messages = s.db.C("messages")
	s.addNeededIndexes()
	return &s, nil
}

func (m *Mongo) addNeededIndexes() {
	m.messages.EnsureIndex(mgo.Index{
		Key:        []string{"timestamp"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     false,
	})
}
