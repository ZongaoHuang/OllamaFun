// Package main 提供了基于Eino库实现的Ollama聊天功能示例
package main

import (
	"context"
	"log"
	"time"
)

// main 函数是程序的入口点，实现了基于Ollama模型的聊天功能流程
func main() {
	// 创建上下文，用于整个应用程序的生命周期管理
	ctx := context.Background()
	log.Println("==create Message==")
	// 从模板创建聊天消息
	messages := createMessageFromTemplate()
	log.Printf("messages: %+v\n\n", messages)

	log.Printf("===create llm===\n")
	// 创建Ollama聊天模型实例
	llm := createOllamaChatModel(ctx)
	log.Printf("create llm success\n\n")

	log.Printf("===llm generate===\n")
	// 创建带超时的上下文，确保生成过程不会无限期运行
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	// 使用模型生成回复
	result := generate(ctxWithTimeout, llm, messages)
	log.Printf("result: %+v\n\n", result)

	log.Printf("===llm stream===\n")
	// 使用流式处理方式获取模型回复
	streamResult := stream(ctx, llm, messages)
	// 输出流式处理结果
	reportStream(streamResult)
}
