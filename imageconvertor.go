package main

import (
	"github.com/icando/chatgpt-proxy-server/kitex_gen/api"
	"github.com/sashabaranov/go-openai"
)

func convertImageRequest(rawReq *api.ImageRequest) openai.ImageRequest {
	req := openai.ImageRequest{}

	if rawReq != nil {
		req.Prompt = rawReq.Prompt
		req.N = int(rawReq.N)
		req.Size = rawReq.Size
		req.ResponseFormat = rawReq.ResponseFormat
		req.User = rawReq.User
	}

	return req
}

func convertImageResponse(rawRsp openai.ImageResponse) *api.ImageResponse {
	return &api.ImageResponse{
		Created: rawRsp.Created,
		Data:    convertImageResponseDataInner(rawRsp.Data),
	}
}

func convertImageResponseDataInner(data []openai.ImageResponseDataInner) []*api.ImageResponseDataInner {
	var r []*api.ImageResponseDataInner
	for _, d := range data {
		r = append(r, &api.ImageResponseDataInner{
			URL:     d.URL,
			B64JSON: d.B64JSON,
		})
	}
	return r
}
