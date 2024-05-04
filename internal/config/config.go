package config

const (
	token = "token"
)

type Config struct {
	TOKEN string
}

func NewConfig() *Config {
	return &Config{
		TOKEN: token,
	}
}
