package config

import (
	"github.com/7thFox/hypothesisbot/src/log"
	"github.com/bwmarrin/discordgo"
)

func (c *Config) Logger(session *discordgo.Session) log.Logger {
	if c.lgr == nil {
		c.lgr = log.NewMultiLogger()

		if c.logConsole() {
			c.lgr.Attach(log.NewConsoleLogger())
		}
		if n := c.logDbName(); n != "" {
			// TODO
		}
	}

	// Allows attaching the session later in the cycle
	if cid := c.logChannelID(); cid != "" && session != nil && !c.lgrHasSession {
		c.lgr.Attach(log.NewChannelLogger(session, cid))
		c.lgrHasSession = true
	}

	return c.lgr
}

func (c *Config) logConsole() bool {
	if c.Debug {
		return c.json.Debug.Logging.Console
	}
	return c.json.Global.Logging.Console
}

func (c *Config) logDbName() string {
	if c.Debug && c.json.Debug.Logging.Dbname != "" {
		return c.json.Debug.Logging.Dbname
	}
	return c.json.Global.Logging.Dbname
}

func (c *Config) logChannelID() string {
	if c.Debug && c.json.Debug.Logging.ChannelID != "" {
		return c.json.Debug.Logging.ChannelID
	}
	return c.json.Global.Logging.ChannelID
}
