package config

func (c *Config) Prefix() string {
	if c.Debug && c.json.Debug.Prefix != "" {
		return c.json.Debug.Prefix
	}
	return c.json.Global.Prefix
}
