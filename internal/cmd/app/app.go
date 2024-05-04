package app

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"log"
	"postmanbot/internal/config"
	"postmanbot/internal/middleware"
	"postmanbot/internal/service/handlers"
)

var (
	cfg    = config.NewConfig()
	logger = middleware.NewLogger()
	shg    = handlers.NewStaticGroupHandler()
	dhg    = handlers.NewDynamicGroupHandler()
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

	bh.Handle(shg.HandleStart, th.CommandEqual("start")) // реагирует на команду /start
	bh.Handle(shg.HandleBack, th.TextContains("Назад"))
	bh.Handle(shg.HandleFAQ, th.TextContains("FAQ"))
	bh.Handle(dhg.HandleTrackPkg, th.TextContains("Отследить посылку"))
	bh.Handle(dhg.HandleRemindPkg, th.TextContains("Добавить напоминание"))
	bh.Start()
}
