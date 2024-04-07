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

func GetCommands() []Command {
	return commands
}
