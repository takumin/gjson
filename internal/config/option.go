package config

type Option interface {
	Apply(*Config)
}

type LogLevel string

func (o LogLevel) Apply(c *Config) {
	c.LogLevel = string(o)
}

type Variable string

func (o Variable) Apply(c *Config) {
	c.Variable = string(o)
}
