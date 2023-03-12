# chatgpt-proxy-server
基于 kitex 的 ChatGPT 代理服务。
![chatgpt-proxy-server.png.png](./docs/images/chatgpt-proxy-server.png)

## 调用样例

```
func main() {
	// ip地址和端口改成 chatgpt-proxy-server 部署的地址和端口
	client, err := gptservice.NewClient("xxx.yyy.chatgpt", client.WithHostPorts("0.0.0.0:8888"))
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
```

## 详细文档

### 服务端
* 配置api key。在/conf/config.prod.yaml下配置：

```
# GPT 相关配置
gptConfig:
  api_key: your_api_key
```

* 服务运行 
```
sh script/start_prod.sh
```

### 客户端

详见client/示例

* 配置服务端地址

```
client, err := gptservice.NewClient("xxx.yyy.chatgpt", client.WithHostPorts("0.0.0.0:8888"))
```

* 执行

    执行client/main.go


## 其他

OpenAI加大了对国内的封锁，国内访问ChatGPT越来越难。

ChatGPT很强大，我搭建了一个服务方便学习研究，免注册使用，欢迎大家一起交流体验~

![欢迎交流](./docs/images/扫码_搜索联合传播样式-标准色版.jpg)