package config

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	cfg *Config
)

//Config example config
type Config struct {
	LogPath   string     `yaml:"logPath"`
	GtpConfig *GtpConfig `yaml:"gptConfig"`
}

// GtpConfig 项目配置
type GtpConfig struct {
	// gtp apikey
	ApiKey string `yaml:"api_key"`
}

//GetConfig 获取配置
func GetConfig() *Config {
	if cfg != nil {
		return cfg
	}
	configPath := "./conf/config.yaml"
	klog.Info("load config: " + configPath)
	cfgFile := flag.String("config", configPath, "配置文件路径")

	bytes, err := ioutil.ReadFile(*cfgFile)
	if err != nil {
		panic(err)
	}

	cfgData := &Config{}
	err = yaml.Unmarshal(bytes, cfgData)
	if err != nil {
		panic(err)
	}
	cfg = cfgData
	return cfg
}
