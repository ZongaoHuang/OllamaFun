package main

import (
	"context"
	"log"
	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/components/model"
)

func createOllamaChatModel(ctx context.Context) model.ToolCallingChatModel{
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://222.20.126.72:11434",
		Model:"llama3.2",
	})
	if err != nil{
		log.Fatalf("create chat model error: %v\n", err)
	}
	return chatModel
}

// func main(){
// 	ctx := context.Background()
// 	chatModel := createOllamaChatModel(ctx)
// 	message := []*schema.Message{
// 		{
// 			Role:    schema.System,
// 			Content: "你是一个有帮助的助手。",
// 		 },
// 		 {
// 			Role:    schema.User,
// 			Content: "你好！",
// 		 },
// 	}
// 	response, err := chatModel.Generate(ctx, message)
// 	if err != nil{
// 		log.Fatalf("generate error: %v\n", err)
// 	}
// 	fmt.Print(response.Content)
// }