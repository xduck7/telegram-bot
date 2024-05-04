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
	hg     = service.NewHandlerGroup()
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

	bh.Handle(hg.HandleStart, th.CommandEqual("start")) // реагирует на команду /start
	bh.Handle(hg.HandleTrackPkg, th.TextContains("Отследить посылку"))
	bh.Handle(hg.HandleBack, th.TextContains("Назад"))
	bh.Handle(hg.HandleFAQ, th.TextContains("FAQ"))
	bh.Handle(hg.HandleAny, th.AnyMessage()) // не знает команду

	bh.Start()
}
