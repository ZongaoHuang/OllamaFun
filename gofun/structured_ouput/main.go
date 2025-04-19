package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"fmt"
	"github.com/ollama/ollama/api"
)

type CountryInfo struct {
	Capital string `json:"capital"`
	Population float64 `json:"population"`
	Area float64 `json:"area"`
}

func main(){
	client, err := api.ClientFromEnvironment()
	if err != nil{
		log.Fatal(err)
	}

	mseeages := []api.Message{
		api.Message{
			Role:    "user",
			Content: "请介绍美国的首都，人口和面积，并以Json格式返回",
		},
	}

	ctx := context.Background()
	req := &api.ChatRequest{
		Model:    "qwen2.5:72b",
		Messages: mseeages,
		Stream:   new(bool),
		Format:  []byte(`"json"`),
		Options: map[string]interface{}{
			"temperature": 0.0,
		},
	}

	respFunc := func(resp api.ChatResponse) error{
		fmt.Printf("%s\n", strings.TrimSpace(resp.Message.Content))
		var info CountryInfo
		err := json.Unmarshal([]byte(resp.Message.Content), &info)
		if err != nil{
			log.Fatal(err) 
		}
		fmt.Printf("首都：%s\n", info.Capital)
		fmt.Printf("人口：%f\n", info.Population)
		fmt.Printf("面积：%f\n", info.Area)
		return nil
	}

	err = client.Chat(ctx, req, respFunc)
	if err != nil{
		log.Fatal(err)
	}
}