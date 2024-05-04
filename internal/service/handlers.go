package service

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
)

type handlerGroup interface {
	HandleStart(bot *telego.Bot, update telego.Update)
	HandleAny(bot *telego.Bot, update telego.Update)
	HandleBack(bot *telego.Bot, update telego.Update)
	HandleTrackPkg(bot *telego.Bot, update telego.Update)
	HandleFAQ(bot *telego.Bot, update telego.Update)
}

type HandlerGroup struct{}

func NewHandlerGroup() *HandlerGroup {
	return &HandlerGroup{}
}

func (hg *HandlerGroup) HandleStart(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := GetKeyboard("Default")
	keyboard.ResizeKeyboard = true
	message := tu.Message(
		chatID,
		"Привет!\nЯ бот для отслеживания посылок",
	).WithReplyMarkup(keyboard)

	_, err := bot.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}

func (hg *HandlerGroup) HandleBack(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := GetKeyboard("Default")
	keyboard.ResizeKeyboard = true
	message := tu.Message(
		chatID,
		"Меню",
	).WithReplyMarkup(keyboard)

	_, err := bot.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}

func (hg *HandlerGroup) HandleTrackPkg(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := GetKeyboard("Cancel")
	keyboard.ResizeKeyboard = true
	message := tu.Message(
		chatID,
		"Введите ваш трек-номер",
	).WithReplyMarkup(keyboard)

	_, err := bot.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}

func (hg *HandlerGroup) HandleAny(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := GetKeyboard("Default")
	keyboard.ResizeKeyboard = true
	message := tu.Message(
		chatID,
		"К сожалению, я не знаю такие команды ",
	).WithReplyMarkup(keyboard)

	_, err := bot.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}

func (hg *HandlerGroup) HandleFAQ(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := GetKeyboard("Default")
	keyboard.ResizeKeyboard = true
	message := tu.Message(
		chatID,
		"FAQ",
	).WithReplyMarkup(keyboard)

	_, err := bot.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}
