package utils

import (
	"context"
	"fmt"
	"os"

	openai_api "github.com/sashabaranov/go-openai"
)

var OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")

func buildPrompt(inputText string) string {
	return fmt.Sprintf(`
			You are in charge of creating the social media for a popular account.
			You are given longer texts and expected to create a short form version of it suitable for posting to Twitter.
			This should be no longer than 240 characters, include the tags of any other accounts mentioned, and add relevant hashtags.
			---
			Original text:
			%s
			---
		`, inputText)
}

func GenerateTweet(inputText string) (string, error) {
	if OPENAI_API_KEY == "" {
		fmt.Println("OPENAI_API_KEY is required")
		os.Exit(1)
	}
	prompt := buildPrompt(inputText)
	client := openai_api.NewClient(OPENAI_API_KEY)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai_api.ChatCompletionRequest{
			Model: openai_api.GPT3Dot5Turbo,
			Messages: []openai_api.ChatCompletionMessage{
				{
					Role:    openai_api.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
