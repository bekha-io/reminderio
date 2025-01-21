package main

import (
	"context"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var client *openai.Client

func InitClient() {
	cl := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)
	client = cl
}

func generateText(ctx context.Context, prompt string) (string, error) {
	cc, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: openai.F(openai.ChatModelGPT4oMini),
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(prompt),
		}),
		MaxCompletionTokens: openai.Int(75),
		FrequencyPenalty: openai.Float(2),
	})
	if err != nil {
		return "", err
	}
	return cc.Choices[0].Message.Content, nil
}
