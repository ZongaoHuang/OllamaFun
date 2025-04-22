// Package main 提供了实现MCP (Model Context Protocol) 服务器的示例
// 该服务器使用Ollama API处理生成请求，并通过HTTP接口提供服务
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ollama/ollama/api"
)

// MCPRequest 定义了MCP请求的结构
type MCPRequest struct {
	Context string `json:"context"` // 系统提示或上下文信息
	Input   string `json:"input"`   // 用户输入内容
}

// MCPResponse 定义了MCP响应的结构
type MCPResponse struct {
	Output string `json:"output"` // 模型生成的输出内容
}

// main 函数是程序的入口点，设置并启动MCP HTTP服务器
func main() {
	// 设置HTTP服务器
	http.HandleFunc("/generate", handleGenerate)

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认端口为8080
	}
	fmt.Printf("MCP服务器运行在端口 %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}

// handleGenerate 处理/generate端点的HTTP请求
// 接收JSON格式的请求，调用Ollama模型，并返回生成结果
func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只接受POST请求", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求数据
	var req MCPRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "请求格式不正确: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 调用Ollama API处理请求
	output, err := generateWithOllama(req.Context, req.Input)
	if err != nil {
		http.Error(w, "模型处理失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回结果
	w.Header().Set("Content-Type", "application/json")
	resp := MCPResponse{Output: output}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "编码响应失败: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// generateWithOllama 使用Ollama API生成文本
// 接收上下文和用户输入，调用Ollama API，并返回生成的文本
func generateWithOllama(contextStr, input string) (string, error) {
	// 创建Ollama客户端
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return "", fmt.Errorf("创建Ollama客户端失败: %w", err)
	}

	// 构造消息
	messages := []api.Message{
		{
			Role:    "system",
			Content: contextStr, // 系统角色消息作为上下文
		},
		{
			Role:    "user",
			Content: input, // 用户输入作为用户角色消息
		},
	}

	// 设置请求
	req := &api.ChatRequest{
		Model:    "qwen2.5:72b", // 使用qwen2.5:72b模型
		Messages: messages,
		Options: map[string]interface{}{
			"temperature": 0.7, // 设置温度为0.7，保持一定的创造性
		},
	}

	ctx := context.Background()
	responseText := ""

	// 处理响应
	respFunc := func(resp api.ChatResponse) error {
		responseText = resp.Message.Content
		return nil
	}

	// 发送请求
	if err := client.Chat(ctx, req, respFunc); err != nil {
		return "", fmt.Errorf("调用Ollama API失败: %w", err)
	}

	return responseText, nil
}
