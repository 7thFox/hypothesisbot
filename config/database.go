package config

import (
	"fmt"

	"github.com/7thFox/hypothesisbot/database"
)

func (c *Config) Database() database.Database {
	if c.db == nil {
		db, err := database.NewDatabase(c.dbType(), c.dbLocation())
		if err != nil {
			fmt.Println(err.Error())
		}
		c.db = db.(database.Database)
	}
	return c.db
}

func (c *Config) dbType() string {
	if c.Debug && c.json.Debug.Database.Dbtype != "" {
		return c.json.Debug.Database.Dbtype
	}
	return c.json.Global.Database.Dbtype
}
func (c *Config) dbLocation() string {
	if c.Debug && c.json.Debug.Database.Location != "" {
		return c.json.Debug.Database.Location
	}
	return c.json.Global.Database.Location
}
