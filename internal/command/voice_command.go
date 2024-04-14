package command

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"path/filepath"
	"strings"
	"whisper-bot/internal/service"
)

type VoiceCommand struct {
	Bot *tgbotapi.BotAPI
	Client *service.OpenAIClient
}

func (c VoiceCommand) CanRun(update tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Voice != nil
}

func (c VoiceCommand) Run(update tgbotapi.Update) {
	log.Printf("%+v\n", update.Message.Voice)

	url, err := c.Bot.GetFileDirectURL(update.Message.Voice.FileID)

	if err != nil {
		log.Println(err)
		return
	}

	filePath, err := service.DownloadFileFromURL(url)

	if err != nil {
		log.Println(err)
		return
	}

	filePathWithoutExtension := strings.TrimSuffix(filePath, filepath.Ext(filePath))
	outputFilePath := fmt.Sprintf("%s.%s", filePathWithoutExtension, "mp3")

	err = service.ConvertOGGtoMP3(filePath, outputFilePath)

	if err != nil {
		log.Println(err)
		return
	}

	text, err := c.Client.ConvertSpeechToText(outputFilePath)

	if err != nil {
		log.Println(err)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	_, err = c.Bot.Send(msg)

	log.Printf("%+v\n", err )
}
