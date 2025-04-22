// Package main 提供了使用Ollama API获取结构化JSON输出的示例
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ollama/ollama/api"
)

// CountryInfo 定义了国家信息的结构体，用于解析模型返回的JSON数据
type CountryInfo struct {
	Capital    string  `json:"capital"`    // 首都
	Population float64 `json:"population"` // 人口数量
	Area       float64 `json:"area"`       // 国土面积
}

// main 函数是程序的入口点，演示了如何使用Ollama API获取结构化的JSON格式输出
func main() {
	// 从环境变量创建Ollama客户端
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	// 构造用户请求消息
	mseeages := []api.Message{
		api.Message{
			Role:    "user",
			Content: "请介绍美国的首都，人口和面积，并以Json格式返回",
		},
	}

	// 创建上下文
	ctx := context.Background()
	// 构造聊天请求，指定使用qwen2.5:72b模型，并要求返回JSON格式
	req := &api.ChatRequest{
		Model:    "qwen2.5:72b",
		Messages: mseeages,
		Stream:   new(bool),
		Format:   []byte(`"json"`), // 指定返回JSON格式
		Options: map[string]interface{}{
			"temperature": 0.0, // 设置温度为0，使输出更加确定性
		},
	}

	// 定义响应处理函数
	respFunc := func(resp api.ChatResponse) error {
		fmt.Printf("%s\n", strings.TrimSpace(resp.Message.Content))
		// 解析JSON响应到CountryInfo结构体
		var info CountryInfo
		err := json.Unmarshal([]byte(resp.Message.Content), &info)
		if err != nil {
			log.Fatal(err)
		}
		// 打印解析后的结构化数据
		fmt.Printf("首都：%s\n", info.Capital)
		fmt.Printf("人口：%f\n", info.Population)
		fmt.Printf("面积：%f\n", info.Area)
		return nil
	}

	// 发送聊天请求并处理响应
	err = client.Chat(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}
