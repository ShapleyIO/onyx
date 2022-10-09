package config

type Config struct {
	SomeFlag string
}

func LoadDefaultConfig() *Config {
	return &Config{}
}
