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
	Context string `json:"context"`
	Input   string `json:"input"`
}

// MCPResponse 定义了MCP响应的结构
type MCPResponse struct {
	Output string `json:"output"`
}

func main() {
	// 设置HTTP服务器
	http.HandleFunc("/generate", handleGenerate)

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("MCP服务器运行在端口 %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}

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

func generateWithOllama(contextStr, input string) (string, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return "", fmt.Errorf("创建Ollama客户端失败: %w", err)
	}

	// 构造消息
	messages := []api.Message{
		{
			Role:    "system",
			Content: contextStr,
		},
		{
			Role:    "user",
			Content: input,
		},
	}

	// 设置请求
	req := &api.ChatRequest{
		Model:    "qwen2.5:72b", // 可以设置为默认模型或环境变量
		Messages: messages,
		Options: map[string]interface{}{
			"temperature": 0.7,
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