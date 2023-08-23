package config

type Option interface {
	Apply(*Config)
}

type LogLevel string

func (o LogLevel) Apply(c *Config) {
	c.LogLevel = string(o)
}

type SearchPath string

func (o SearchPath) Apply(c *Config) {
	c.Path.Searches = append(c.Path.Searches, string(o))
}

type Includes string

func (o Includes) Apply(c *Config) {
	c.Extention.Includes = append(c.Extention.Includes, string(o))
}

type Excludes string

func (o Excludes) Apply(c *Config) {
	c.Extention.Excludes = append(c.Extention.Excludes, string(o))
}
