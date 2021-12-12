package config

// Config is struct for Config
type Config struct {
	Server Server
}

// NewConfig returns a new instance of Config instance
func NewConfig() *Config {
	return &Config{}
}
