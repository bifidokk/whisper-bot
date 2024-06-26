package service

import (
	"context"
	"log"
	"whisper-bot/internal/config"

	"github.com/sashabaranov/go-openai"
)

const prompt string = "Use capital letters and punctuation. Do not repeat yourself. Do not describe ambient sounds or noise or silence, just ommit. Break the text into paragraphs. Separate paragraphs with blank lines"

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient() *OpenAIClient {
	authToken := config.GetEnv("openai_token")

	client := openai.NewClient(authToken)

	return &OpenAIClient{
		client: client,
	}
}

func (c OpenAIClient) ConvertSpeechToText(inputFile string) (string, error) {
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: inputFile,
		Prompt:   prompt,
	}

	resp, err := c.client.CreateTranscription(ctx, req)

	if err != nil {
		log.Printf("Transcription error: %v\n", err)
		return "", err
	}

	log.Println(resp.Text)

	return resp.Text, nil
}
