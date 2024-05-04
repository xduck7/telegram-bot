package handlers

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"postmanbot/internal/service"
)

type StaticHandlerGroup interface {
	HandleStart(bot *telego.Bot, update telego.Update)
	HandleBack(bot *telego.Bot, update telego.Update)
	HandleFAQ(bot *telego.Bot, update telego.Update)
}

type staticHandlerGroup struct {
	StaticHandlerGroup
}

func NewStaticGroupHandler() StaticHandlerGroup {
	return &staticHandlerGroup{}
}

func (sgh *staticHandlerGroup) HandleStart(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := service.GetKeyboard("Default")
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

func (sgh *staticHandlerGroup) HandleBack(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := service.GetKeyboard("Default")
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

func (sgh *staticHandlerGroup) HandleFAQ(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := service.GetKeyboard("Default")
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
