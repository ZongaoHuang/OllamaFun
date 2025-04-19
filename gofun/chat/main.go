package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

func main(){
	client ,err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	message := []api.Message{
		api.Message{
			Role:    "user",
			Content: "武汉今天的天气一般怎样？",
		},
	}

	ctx := context.Background()
	req := & api.ChatRequest{
		Model:   "qwen2.5:72b",
		Messages: message,
	}

	respFunc := func(resp api.ChatResponse) error{
		fmt.Print(resp.Message.Content)
		return nil
	}

	err = client.Chat(ctx, req, respFunc)

	if err != nil {
		log.Fatal(err)
	}
}