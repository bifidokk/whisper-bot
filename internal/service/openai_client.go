package service

import (
	"context"
	"log"
	"whisper-bot/internal/config"

	"github.com/sashabaranov/go-openai"
)

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

func (c OpenAIClient) UploadFile(inputFile string) {
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: inputFile,
	}

	resp, err := c.client.CreateTranscription(ctx, req)

	if err != nil {
		log.Printf("Transcription error: %v\n", err)
		return
	}

	log.Println(resp.Text)
}
