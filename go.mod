module github.com/icando/chatgpt-proxy-server

go 1.16

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/thrift v0.13.0
	github.com/cloudwego/kitex v0.5.0
)

require (
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/sashabaranov/go-openai v1.5.7
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
	gopkg.in/yaml.v3 v3.0.1
)
