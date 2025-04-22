// Package main 提供了使用Ollama官方API进行简单聊天的示例
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

// main 函数是程序的入口点，演示了如何使用Ollama API进行基本的聊天交互
func main() {
	// 从环境变量创建Ollama客户端
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	// 构造用户消息
	message := []api.Message{
		api.Message{
			Role:    "user",
			Content: "武汉今天的天气一般怎样？",
		},
	}

	// 创建上下文
	ctx := context.Background()

	// 构造聊天请求，指定使用qwen2.5:72b模型
	req := &api.ChatRequest{
		Model:    "qwen2.5:72b",
		Messages: message,
	}

	// 定义响应处理函数，直接打印模型返回的内容
	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		return nil
	}

	// 发送聊天请求并处理响应
	err = client.Chat(ctx, req, respFunc)

	if err != nil {
		log.Fatal(err)
	}
}
