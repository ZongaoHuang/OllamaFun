package main

import(
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main(){
	llm, err := ollama.New(
		ollama.WithModel("llama3.2"),
		ollama.WithServerURL("http://222.20.126.72:11434"),
	)
	if err != nil{
		log.Fatal(err)
	}

	ctx := context.Background()
	completion, err := llm.Call(ctx, "Go语言的特点是什么？",
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte)error{
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil{
		log.Fatal(err)
	}
	_ = completion
	fmt.Println("完成：", completion)
}