package command

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type VoiceCommand struct {
	Bot *tgbotapi.BotAPI
}

func (c VoiceCommand) CanRun(update tgbotapi.Update) bool {
	return true
}

func (c VoiceCommand) Run(update tgbotapi.Update) {
	fmt.Printf("%+v\n", update.Message)
}
