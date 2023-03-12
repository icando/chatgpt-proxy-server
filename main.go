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
