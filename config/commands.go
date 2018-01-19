package config

import (
	"github.com/7thFox/hypothesisbot/command"
)

func (c *Config) Commands() map[string]command.Command {
	if c.cmds == nil {
		cmdsAll := []command.Command{
			command.NewDoot(),
			command.NewGit(),
			command.NewHoot(),
			command.NewKill(),
			command.NewNoot(),
			command.NewSay(),
			command.NewTest(),
			command.NewVersion(c.Version),
		}

		cmdsAll = append(cmdsAll, command.NewHelp(cmdsAll))

		// read config and construct commands
		c.cmds = map[string]command.Command{}
		for _, cmd := range cmdsAll {
			if !c.blacklist()[cmd.Name()] {
				c.cmds[cmd.Name()] = cmd
			}
		}
	}
	return c.cmds
}

func (c *Config) blacklistArray() []string {

	if c.Debug && c.json.Debug.CmdBlacklist != nil {
		return c.json.Debug.CmdBlacklist
	}
	return c.json.Global.CmdBlacklist
}

func (c *Config) blacklist() map[string]bool {
	if c.blacklistMap == nil {
		c.blacklistMap = map[string]bool{}
		for _, b := range c.blacklistArray() {
			c.blacklistMap[b] = true
		}
	}
	return c.blacklistMap
}
