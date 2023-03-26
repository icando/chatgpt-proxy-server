package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/cloudwego/kitex/client"
	"github.com/icando/chatgpt-proxy-server/kitex_gen/api"
	"github.com/icando/chatgpt-proxy-server/kitex_gen/api/gptservice"
	"github.com/sashabaranov/go-openai"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestCreateImage(t *testing.T) {
	client, err := gptservice.NewClient("chatgpt.proxy.server", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.ImageRequest{
		Prompt:         "Parrot on a skateboard performs a trick, cartoon style, natural light, high detail",
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}
	respUrl, err := client.CreateImage(context.Background(), req)
	if err != nil {
		log.Printf("Image creation error: %v\n", err)
		return
	}
	log.Println(respUrl.Data[0].URL)
}

func TestCreateImage_Base64(t *testing.T) {
	client, err := gptservice.NewClient("chatgpt.proxy.server", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	// Example image as base64
	reqBase64 := &api.ImageRequest{
		Prompt:         "Portrait of a humanoid parrot in a classic costume, high detail, realistic light, unreal engine",
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := client.CreateImage(context.Background(), reqBase64)
	if err != nil {
		log.Printf("Image creation error: %v\n", err)
		return
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		log.Printf("Base64 decode error: %v\n", err)
		return
	}

	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		log.Printf("PNG decode error: %v\n", err)
		return
	}

	file, err := os.Create("example.png")
	if err != nil {
		log.Printf("File creation error: %v\n", err)
		return
	}
	defer file.Close()

	if err := png.Encode(file, imgData); err != nil {
		log.Printf("PNG encode error: %v\n", err)
		return
	}

	log.Println("The image was saved as example.png")
}
