package config

import (
	"fmt"
	"os/exec"

	"github.com/7thFox/hypothesisbot/database"
)

func (c *Config) Database() database.Database {
	if c.db == nil {
		db, err := database.NewDatabase(c.dbType(), c.dbHost(), c.dbName())
		if err != nil {
			c.Logger(nil).Log(fmt.Sprintf("Error getting database: %s", err.Error()))
			c.Logger(nil).Log(fmt.Sprintf("Host: %s, Name: %s Type: %s", c.dbHost(), c.dbName(), c.dbType()))
			out, _ := exec.Command("ping", c.dbHost(), "-c 5", "-w 10").Output()
			c.Logger(nil).Log(string(out))
			return nil
		}
		c.Logger(nil).Log("DB Connected")
		c.db = db.(database.Database)
	}
	return c.db
}

func (c *Config) dbName() string {
	if c.Debug && c.json.Debug.Database.Dbname != "" {
		return c.json.Debug.Database.Dbname
	}
	return c.json.Global.Database.Dbname
}

func (c *Config) dbType() string {
	if c.Debug && c.json.Debug.Database.Dbtype != "" {
		return c.json.Debug.Database.Dbtype
	}
	return c.json.Global.Database.Dbtype
}
func (c *Config) dbHost() string {
	if c.Debug && c.json.Debug.Database.Host != "" {
		return c.json.Debug.Database.Host
	}
	return c.json.Global.Database.Host
}
