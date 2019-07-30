package botcontrol

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotInterface interface {
	RunBot() *tgbotapi.BotAPI
	MakeNewUpdate() tgbotapi.UpdateConfig
}

func RunBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func MakeNewUpdate() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	return updateConfig
}
