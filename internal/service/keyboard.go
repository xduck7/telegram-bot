package service

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func GetKeyboard(action string) *telego.ReplyKeyboardMarkup {
	switch action {
	case "Cancel":
		return tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("Назад🔙"),
			),
		)

	default:
		return tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("Отследить посылку📍"),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("Добавить напоминание📅"),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("FAQ📌"),
			),
		)

	}
}
