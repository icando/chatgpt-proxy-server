// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
	client, err := gptservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.ChatCompletionRequest{Model: openai.GPT3Dot5Turbo,
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
