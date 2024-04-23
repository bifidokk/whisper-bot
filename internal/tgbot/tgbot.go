package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"whisper-bot/internal/config"
)

var webhookPath string

func Create() *tgbotapi.BotAPI {
	botToken := config.GetEnv("token")
	webhookUrl := config.GetEnv("webhook_url")
	webhookPath = config.GetEnv("webhook_path")
	baseUrl := config.GetEnv("base_url")

	go func() {
		err := http.ListenAndServe(baseUrl, nil)
		if err != nil {
			panic(err)
		}
	}()

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookUrl + webhookPath + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	return bot
}

func GetUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	return bot.ListenForWebhook(webhookPath + bot.Token)
}
