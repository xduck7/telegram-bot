package app

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"log"
	"postmanbot/internal/config"
	"postmanbot/internal/middleware"
	"postmanbot/internal/service"
)

var (
	cfg    = config.NewConfig()
	logger = middleware.NewLogger()
)

func Run() {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from", r)
		}
	}()

	bot, err := telego.NewBot(cfg.TOKEN, logger)
	if err != nil {
		log.Fatal(err)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, err := th.NewBotHandler(bot, updates)

	defer bot.StopLongPolling()
	defer bh.Stop()

	bh.Handle(service.HandleStart, th.CommandEqual("start"))
	bh.Handle(service.HandleBack, th.CommandEqual("Назад"))
	bh.Handle(service.HandleTrackPkg, th.CommandEqual("Отследить посылку"))

	bh.Start()
}
