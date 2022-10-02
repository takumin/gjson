package config

type Config struct {
	LogLevel string

	Path struct {
		File      string
		Directory string
	}

	Extention struct {
		Includes string
		Excludes string
	}
}

func NewConfig(opts ...Option) *Config {
	c := &Config{}
	for _, o := range opts {
		o.Apply(c)
	}
	return c
}
