package config

func (c *Config) LogServers() []string {
	if c.Debug && c.json.Debug.LogServers != nil {
		return c.json.Debug.LogServers
	}
	return c.json.Global.LogServers
}
