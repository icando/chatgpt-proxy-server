package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/icando/chatgpt-proxy-server/config"
	"github.com/sashabaranov/go-openai"
	"io"
	"os"

	api "github.com/icando/chatgpt-proxy-server/kitex_gen/api/gptservice"
)

func main() {
	cfg := config.GetConfig()
	logPath := cfg.LogPath
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		klog.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	mw := io.MultiWriter(f, os.Stdout)
	klog.SetOutput(mw)

	svr := api.NewServer(&GPTServiceImpl{
		client: openai.NewClient(cfg.GtpConfig.ApiKey),
	})

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}
