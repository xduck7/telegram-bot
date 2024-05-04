package config

const (
	token = "TOKEN"
)

type Config struct {
	TOKEN string
}

func NewConfig() *Config {
	return &Config{
		TOKEN: token,
	}
}
