package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"whisper-bot/internal/config"
)

func main() {
	botToken := config.GetEnv("token")
	webhookUrl := config.GetEnv("webhook_url")
	webhookPath := config.GetEnv("webhook_path")
	baseUrl := config.GetEnv("base_url")

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

	updates := bot.ListenForWebhook(webhookPath + bot.Token)

	go func() {
		err := http.ListenAndServe(baseUrl, nil)
		if err != nil {
			panic(err)
		}
	}()

	log.Println("Listening on " + webhookUrl)

	for update := range updates {
		if update.Message == nil {

		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}
