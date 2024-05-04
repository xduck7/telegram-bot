package handlers

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"postmanbot/internal/service"
	"strings"
	"time"
)

type DynamicHandlerGroup interface {
	HandleTrackPkg(bot *telego.Bot, update telego.Update)
	HandleRemindPkg(bot *telego.Bot, update telego.Update)
}

type dynamicHandlerGroup struct {
	DynamicHandlerGroup
}

func NewDynamicGroupHandler() DynamicHandlerGroup {
	return &dynamicHandlerGroup{}
}

func (dhg *dynamicHandlerGroup) HandleTrackPkg(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := service.GetKeyboard("Cancel")
	keyboard.ResizeKeyboard = true

	message := tu.Message(
		chatID,
		"Введите ваш трек-номер",
	).WithReplyMarkup(keyboard)

	_, err := bot.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}

	updateChan := make(chan telego.Update)

	go func() {
		for {
			updates, err := bot.GetUpdates(nil)

			if err != nil {
				log.Fatal(err)
			}
			for _, update := range updates {
				// Отправляем обновление в канал
				updateChan <- update
			}
		}
	}()

	var trackNumber string
	select {
	case update := <-updateChan:
		if update.Message == nil {
			log.Fatal("Не получено сообщение")
		}
		trackNumber = update.Message.Text
	case <-time.After(30 * time.Second):
		log.Fatal("Таймаут ожидания ввода трек-номера")
	}

	if !strings.Contains(trackNumber, "Назад") {
		// TODO: функция получения информации о треке
		keyboard = service.GetKeyboard("Default")
		keyboard.ResizeKeyboard = true
		message = tu.Message(
			chatID,
			trackNumber+" отслежен",
		).WithReplyMarkup(keyboard)

		_, err = bot.SendMessage(message)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (dhg *dynamicHandlerGroup) HandleRemindPkg(bot *telego.Bot, update telego.Update) {
	chatID := tu.ID(update.Message.Chat.ID)
	keyboard := service.GetKeyboard("Cancel")
	keyboard.ResizeKeyboard = true

	message := tu.Message(
		chatID,
		"Введите трек-номер, который хотите добавить в напоминания",
	).WithReplyMarkup(keyboard)

	_, err := bot.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}

	updateChan := make(chan telego.Update)

	go func() {
		for {
			updates, err := bot.GetUpdates(nil)

			if err != nil {
				log.Fatal(err)
			}
			for _, update := range updates {
				// Отправляем обновление в канал
				updateChan <- update
			}
		}
	}()

	var trackNumber string
	select {
	case update := <-updateChan:
		if update.Message == nil {
			log.Fatal("Не получено сообщение")
		}
		trackNumber = update.Message.Text
	case <-time.After(30 * time.Second):
		log.Fatal("Таймаут ожидания ввода трек-номера")
	}

	if !strings.Contains(trackNumber, "Назад") {
		// TODO: функция добавления трека в напоминания
		keyboard = service.GetKeyboard("Default")
		keyboard.ResizeKeyboard = true
		message = tu.Message(
			chatID,
			trackNumber+" добавлен",
		).WithReplyMarkup(keyboard)

		_, err = bot.SendMessage(message)
		if err != nil {
			log.Fatal(err)
		}
	}
}
