package command

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Command interface {
	CanRun(update tgbotapi.Update) bool
	Run(update tgbotapi.Update)
}

var commands []Command

func Init(telegramBot *tgbotapi.BotAPI) {
	commands = []Command{
		VoiceCommand{Bot: telegramBot},
	}
}

func Handle(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	for _, command := range commands {
		if command.CanRun(update) {
			command.Run(update)
		}
	}
}
