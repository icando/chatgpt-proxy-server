package main

import (
	"context"
	"github.com/icando/chatgpt-proxy-server/config"
	"github.com/icando/chatgpt-proxy-server/kitex_gen/api"
	"github.com/sashabaranov/go-openai"
)

// GPTServiceImpl implements the last service interface defined in the IDL.
type GPTServiceImpl struct {
	client *openai.Client
}

// Echo implements the GPTServiceImpl interface.
func (s *GPTServiceImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	resp = &api.Response{Message: req.Message}
	return
}

// CreateChatCompletion35 implements the GPTServiceImpl interface.
func (s *GPTServiceImpl) CreateChatCompletion35(ctx context.Context, request *api.ChatCompletionRequest) (resp *api.ChatCompletionResponse, err error) {
	if config.GetConfig().Env == "dev" {
		return &api.ChatCompletionResponse{
			Model: request.Model,
			Choices: []*api.ChatCompletionChoice{{
				Index: 1,
				Message: &api.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleAssistant,
					Content: request.Messages[len(request.Messages)-1].Content,
				},
			}},
		}, nil
	}

	rawRsp, rawErr := s.client.CreateChatCompletion(
		ctx,
		convertRequest(request),
	)

	if rawErr != nil {
		return convertResponse(rawRsp), rawErr
	}

	return convertResponse(rawRsp), nil
}

func (s *GPTServiceImpl) CreateImage(ctx context.Context, request *api.ImageRequest) (resp *api.ImageResponse, err error) {
	rawRsp, rawErr := s.client.CreateImage(
		ctx,
		convertImageRequest(request),
	)

	if rawErr != nil {
		return convertImageResponse(rawRsp), rawErr
	}

	return convertImageResponse(rawRsp), nil
}
