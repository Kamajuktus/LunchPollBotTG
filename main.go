package main

import (
	"TelegramBot/telegram"
	"log"

	_ "github.com/lib/pq"
	tele "gopkg.in/tucnak/telebot.v2"
)

func main() {
	pref, err := tele.NewBot(tele.Settings{
		Token: "5513424668:AAE6PwznYiUxsxa-wqg11Tt9vMXdnDQY2WA",
		//Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b := telegram.NewBot(pref)
	b.Init()

	b.Bot.Start()
}
