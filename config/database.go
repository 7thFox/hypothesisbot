package config

func (c *Config) DatabaseName() string {
	if c.Debug && c.json.Debug.Database.Dbname != "" {
		return c.json.Debug.Database.Dbname
	}
	return c.json.Global.Database.Dbname
}

func (c *Config) DatabaseType() string {
	if c.Debug && c.json.Debug.Database.Dbtype != "" {
		return c.json.Debug.Database.Dbtype
	}
	return c.json.Global.Database.Dbtype
}

func (c *Config) DatabaseHost() string {
	if c.Debug && c.json.Debug.Database.Host != "" {
		return c.json.Debug.Database.Host
	}
	return c.json.Global.Database.Host
}
