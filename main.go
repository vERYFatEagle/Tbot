// tbot project main.go
package main

import (
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var BananaSum int = 1

func main() {

	bot, err := tgbotapi.NewBotAPI("1969957819:AAH4BcGJlZlZ9r_Q7s-zmSX7d-3vwlOA_qM")
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
		case "\xF0\x9F\x8D\x8C":

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Спасибо за \xF0\x9F\x8D\x8C")
			//msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)

			textSUM := fmt.Sprintf("Теперь у меня %d \xF0\x9F\x8D\x8C", BananaSum)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, textSUM)
			//msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
			BananaSum++

		}

	}
}

func (t *tgbotapi.BotAPI) creatReplyKeyboardMarkup(chatID int64) {
msg := tgbotapi.NewMessage(chatID, "ReplyKeyboardMarkup")
msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
        Keyboard: [][]tgbotapi.KeyboardButton{
tgbotapi.NewKeyboardButtonRow(
tgbotapi.NewKeyboardButton («Button(1,1)»),
tgbotapi.NewKeyboardButtonLocation("location(1,2)"),
tgbotapi.NewKeyboardButtonContact("contact(1,3)"),
            ),
tgbotapi.NewKeyboardButtonRow(
tgbotapi.NewKeyboardButton («Button(2,1)»),
tgbotapi.NewKeyboardButtonLocation("location(2,2)"),
tgbotapi.NewKeyboardButtonContact("contact(2,3)"),


            ),
        },
    }bot.creatReplyKeyboardMarkup(update.Message.Chat.ID int64)
 
bot.creatReplyKeyboardMarkup(update.Message.Chat.ID int64)
}
