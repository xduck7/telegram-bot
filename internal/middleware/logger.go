package middleware

import "github.com/mymmrac/telego"

func NewLogger() telego.BotOption {
	return telego.WithDefaultLogger(true, true)
}
