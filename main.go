// tbot project main.go
package main

import (
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var BananaSum int = 0

func main() {

	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		switch update.Message.Text {
		case "+":

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Спасибо за банан")
			//msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)

			textSUM := fmt.Sprintf("Теперь у меня %d", BananaSum)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, textSUM)
			//msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
			BananaSum++
		}

	}

}
