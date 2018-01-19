package config

import (
	"github.com/7thFox/hypothesisbot/command"
)

func (c *Config) Commands() map[string]command.Command {
	if c.cmds == nil {
		// read config and construct commands
		c.cmds = map[string]command.Command{}
		cmds := []command.Command{
			command.NewDoot(),
			command.NewGit(),
			command.NewHoot(),
			command.NewKill(),
			command.NewNoot(),
			command.NewSay(),
			command.NewTest(),
			command.NewVersion(c.Version),
		}
		for _, cmd := range cmds {
			c.cmds[cmd.Name()] = cmd
		}
	}
	return c.cmds
}
