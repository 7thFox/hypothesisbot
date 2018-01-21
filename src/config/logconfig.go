package config

func (c *Config) LogConsole() bool {
	if c.Debug {
		return c.json.Debug.Logging.Console
	}
	return c.json.Global.Logging.Console
}

func (c *Config) LogDbName() string {
	if c.Debug && c.json.Debug.Logging.Dbname != "" {
		return c.json.Debug.Logging.Dbname
	}
	return c.json.Global.Logging.Dbname
}

func (c *Config) LogChannelID() string {
	if c.Debug && c.json.Debug.Logging.ChannelID != "" {
		return c.json.Debug.Logging.ChannelID
	}
	return c.json.Global.Logging.ChannelID
}
