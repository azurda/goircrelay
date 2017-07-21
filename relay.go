package main

import (
	"crypto/tls"
	"fmt"
	"github.com/thoj/go-ircevent"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

type config struct {
	ircnick    string
	channel    string
	ircserver  string
	tgbot_key  string
	tg_channel int64
}

func main() {
	conf := config{
		ircnick:    "gorelay",
		channel:    "#go-irc-test",
		ircserver:  "irc.freenode.net:7000",
		tgbot_key:  "",
		tg_channel: ,
	}
	bot, err := tgbotapi.NewBotAPI(conf.tgbot_key)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	irccon := irc.IRC(conf.ircnick, "IRCTestSSL")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = false
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(conf.channel) })
	irccon.AddCallback("366", func(e *irc.Event) {})
	irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
		log.Printf("[%s] %s", e.Nick, e.Message())
		msg := fmt.Sprintf("<%s> %s", e.Nick, e.Message())
		bot.Send(tgbotapi.NewMessage(conf.tg_channel, msg))
	})

	err = irccon.Connect(conf.ircserver)
	if err != nil {
		log.Printf("Err %s", err)
		return
	}

	go irccon.Loop()

	for update := range updates {

		if update.Message == nil {
			continue
		}
		if update.Message.Chat == nil {
			continue
		}

		username := update.Message.From.UserName
		if len(username) == 0 {
			username = update.Message.From.FirstName
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := fmt.Sprintf("<%s> %s", update.Message.From.UserName, update.Message.Text)
		irccon.Privmsg(conf.channel, msg)
	}
}
