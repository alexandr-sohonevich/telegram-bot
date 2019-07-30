package main

import (
	"testBot/botcontrol"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot := botcontrol.RunBot()
	config := botcontrol.MakeNewUpdate()

	updates, _ := bot.GetUpdatesChan(config)

	for update := range updates {
		// if update.Message == nil { // ignore any non-Message Updates
		// 	continue
		// }

		if update.Message.Text == "/fuck@sohonevichBot" {
			update.Message.Text = "/fuck"
		}

		if update.Message.Text == "/proval@sohonevichBot" {
			update.Message.Text = "/proval"
		}

		if update.Message.Text == "/zaebal@sohonevichBot" {
			update.Message.Text = "/zaebal"
		}

		switch update.Message.From.UserName {
		case "vladoleg":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Проваленок, завали ебало")
			bot.Send(msg)
		}

		switch update.Message.Text {
		case "/fuck":
			fuckYouAnswer := "Fuck you, " + update.Message.From.FirstName
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fuckYouAnswer)
			bot.Send(msg)
		case "/proval":

			if update.Message.From.UserName == "sasha_dotsenko" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Проваленок зайчик")
				bot.Send(msg)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Проваленок, иди нахуй")
			bot.Send(msg)
		case "/zaebal":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Проваленок, как же ты заебал")
			bot.Send(msg)
		default:
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Че за хуйню ты написал? Я таких команд не знаю. Завали ебало сойдет?")
			//bot.Send(msg)
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Проваленок иди нахуй")

	}
}
