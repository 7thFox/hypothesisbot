package config

func (c *Config) blacklistArray() []string {

	if c.Debug && c.json.Debug.CmdBlacklist != nil {
		return c.json.Debug.CmdBlacklist
	}
	return c.json.Global.CmdBlacklist
}

func (c *Config) CommandBlacklist() map[string]bool {
	if c.blacklistMap == nil {
		c.blacklistMap = map[string]bool{}
		for _, b := range c.blacklistArray() {
			c.blacklistMap[b] = true
		}
	}
	return c.blacklistMap
}
