package main

import (
	"whisper-bot/internal/command"
	"whisper-bot/internal/tgbot"
)

func main() {
	bot := tgbot.Create()
	command.Init(bot)

	for update := range tgbot.GetUpdates(bot) {
		command.Handle(update)
	}
}
