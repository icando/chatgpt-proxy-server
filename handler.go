package main

import (
	"context"
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
	rawRsp, rawErr := s.client.CreateChatCompletion(
		ctx,
		convertRequest(request),
	)

	if rawErr != nil {
		return convertResponse(rawRsp), rawErr
	}

	return convertResponse(rawRsp), nil
}

func convertRequest(rawReq *api.ChatCompletionRequest) openai.ChatCompletionRequest {
	req := openai.ChatCompletionRequest{}

	if rawReq != nil {
		req.Model = rawReq.Model
		req.Messages = convertChatCompletionMessage(rawReq.Messages)
		req.MaxTokens = int(rawReq.MaxTokens)
		req.Temperature = float32(rawReq.Temperature)
		req.TopP = float32(rawReq.TopP)
		req.N = int(rawReq.N)
		req.Stream = rawReq.Stream
		req.Stop = rawReq.Stop
		req.PresencePenalty = float32(rawReq.PresencePenalty)
		req.FrequencyPenalty = float32(rawReq.FrequencyPenalty)
		req.LogitBias = convertMapInt32(rawReq.LogitBias)
		req.User = rawReq.User
	}

	return req
}

func convertResponse(rawRsp openai.ChatCompletionResponse) *api.ChatCompletionResponse {
	return &api.ChatCompletionResponse{
		ID:      rawRsp.ID,
		Object:  rawRsp.Object,
		Created: rawRsp.Created,
		Model:   rawRsp.Model,
		Choices: convertChatCompletionChoices(rawRsp.Choices),
		Usage:   convertUsage(rawRsp.Usage),
	}
}

func convertChatCompletionMessage(msg []*api.ChatCompletionMessage) []openai.ChatCompletionMessage {
	var r []openai.ChatCompletionMessage
	for _, m := range msg {
		r = append(r, openai.ChatCompletionMessage{
			Role:    m.Role,
			Content: m.Content,
			Name:    m.Name,
		})
	}
	return r
}

func convertMapInt32(raw map[string]int32) map[string]int {
	r := make(map[string]int)
	for k, v := range raw {
		r[k] = int(v)
	}
	return r
}

func convertChatCompletionChoices(raw []openai.ChatCompletionChoice) []*api.ChatCompletionChoice {
	var r []*api.ChatCompletionChoice
	for _, c := range raw {
		r = append(r, &api.ChatCompletionChoice{
			Index: int32(c.Index),
			Message: &api.ChatCompletionMessage{
				Role:    c.Message.Role,
				Content: c.Message.Content,
				Name:    c.Message.Name,
			},
			FinishReason: c.FinishReason,
		})
	}
	return r
}

func convertUsage(raw openai.Usage) *api.Usage {
	return &api.Usage{
		PromptTokens:     int32(raw.PromptTokens),
		CompletionTokens: int32(raw.CompletionTokens),
		TotalTokens:      int32(raw.TotalTokens),
	}
}
