package main

import (
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	"log"
	"context"
)

func generate(ctx context.Context, llm model.ToolCallingChatModel, in []*schema.Message) *schema.Message{
	result, err := llm.Generate(ctx, in)
	if err != nil{
		log.Fatalf("llm generate error: %v\n", err)
	}
	return result
}

func stream(ctx context.Context, llm model.ToolCallingChatModel, in []*schema.Message) *schema.StreamReader[*schema.Message]{
	result, err := llm.Stream(ctx, in)
	if err != nil{
		log.Fatalf("llm stream error: %v\n", err)
	}
	return result
}