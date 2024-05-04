package service

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
)

//TODO: interface

func HandleStart(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := GetKeyboard("Start")

	message := tu.Message(
		chatID,
		"Привет!\nЯ бот для отслеживания посылок",
	).WithReplyMarkup(keyboard)

	_, err := bot.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleBack(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := GetKeyboard("Start")

	message := tu.Message(
		chatID,
		"втаоптволапол",
	).WithReplyMarkup(keyboard)

	_, err := bot.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleTrackPkg(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := GetKeyboard("Cancel")
	message := tu.Message(
		chatID,
		"Отравьте мне трек-номер для отслеживания",
	).WithReplyMarkup(keyboard)

	_, err := bot.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}
