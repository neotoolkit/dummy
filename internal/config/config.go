package config

type Config struct {
	Server Server
}

func NewConfig() *Config {
	return &Config{}
}
