package mongo

import (
	"gopkg.in/mgo.v2"
)

type Mongo struct {
	session *mgo.Session
}

func NewMongo(host string) (*Mongo, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	s := Mongo{}
	s.session = session
	return &s, nil
}
