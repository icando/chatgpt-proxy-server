include "image.thrift"

namespace go api

struct Request {
	1: string message
}

struct Response {
	1: string message
}

struct ChatCompletionMessage {
	string Role
	string Content

	// This property isn't in the official documentation, but it's in
	// the documentation for the official library for python:
	// - https://github.com/openai/openai-python/blob/main/chatml.md
	// - https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb
	string Name
}

struct ChatCompletionRequest {
    string                  Model
	list<ChatCompletionMessage> Messages
	i32                     MaxTokens
	double                 Temperature
	double                 TopP
	i32                     N
	bool                    Stream
	list<string>               Stop
	double                 PresencePenalty
	double                 FrequencyPenalty
	map<string, i32>          LogitBias
	string                  User
}

struct ChatCompletionChoice {
    i32                   Index
	ChatCompletionMessage Message
	string                FinishReason
}

struct Usage {
	i32 PromptTokens
	i32 CompletionTokens
	i32 TotalTokens
}

struct ChatCompletionResponse {
	string                 ID
	string                 Object
	i64                    Created
	string                 Model
	list<ChatCompletionChoice> Choices
	Usage                  Usage
}

service GPTService {
    Response echo(1: Request req)
    ChatCompletionResponse CreateChatCompletion35(ChatCompletionRequest request)
    image.ImageResponse CreateImage(image.ImageRequest request)
}
