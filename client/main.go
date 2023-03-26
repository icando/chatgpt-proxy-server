package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/icando/chatgpt-proxy-server/kitex_gen/api"
	"github.com/icando/chatgpt-proxy-server/kitex_gen/api/gptservice"
	"github.com/sashabaranov/go-openai"
	"log"
)

func main() {
	client, err := gptservice.NewClient("chatgpt.proxy.server", client.WithHostPorts("1.15.84.181:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []*api.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Hello!",
			},
		},
	}
	resp, err := client.CreateChatCompletion35(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
