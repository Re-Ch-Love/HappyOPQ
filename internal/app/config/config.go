package config

import (
	"HappyOPQ/internal/app/constants"
	"HappyOPQ/pkg/files"
	"HappyOPQ/pkg/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type OPQBotConfig struct {
	ExecPath string `yaml:"ExecPath"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
}

type OneBotHTTPConfig struct {
	Enabled bool   `yaml:"Enabled"`
	URL     string `yaml:"URL"`
}

type OneBotPositiveWebSocketConfig struct {
	Enabled bool   `yaml:"Enabled"`
	Host    string `yaml:"Host"`
	Port    int    `yaml:"Port"`
}

type OneBotReverseWebSocketConfig struct {
	Enabled bool   `yaml:"Enabled"`
	Host    string `yaml:"Host"`
	Port    int    `yaml:"Port"`
}

type OneBotConfig struct {
	HTTP              OneBotHTTPConfig              `yaml:"HTTP"`
	PositiveWebSocket OneBotPositiveWebSocketConfig `yaml:"PositiveWebSocket"`
	ReverseWebSocket  OneBotReverseWebSocketConfig  `yaml:"ReverseWebSocket"`
}

type Config struct {
	OPQBot OPQBotConfig `yaml:"OPQBot"`
	OneBot OneBotConfig `yaml:"OneBot"`
}

var DefaultConfig = Config{
	OPQBot: OPQBotConfig{
		ExecPath: "./OPQBot/OPQBot.exe",
		Host:     "127.0.0.1",
		Port:     8080,
	},
	OneBot: OneBotConfig{
		HTTP: OneBotHTTPConfig{
			Enabled: true,
			URL:     "http://127.0.0.1:8081",
		},
		PositiveWebSocket: OneBotPositiveWebSocketConfig{
			Enabled: false,
			Host:    "127.0.0.1",
			Port:    8082,
		},
		ReverseWebSocket: OneBotReverseWebSocketConfig{
			Enabled: false,
			Host:    "127.0.0.1",
			Port:    8083,
		},
	},
}

func LoadConfig(configPath string) Config {
	if files.IsFileExist(configPath) {
		//configPath = configPath
	} else if files.IsFileExist(constants.DefaultConfigFileName) {
		configPath = constants.DefaultConfigFileName
	} else {
		log.Info("未检测到自定义配置，将使用默认配置")
		out, err := yaml.Marshal(DefaultConfig)
		log.InfoF("当前配置：\n%s", string(out))
		if err != nil {
			log.InfoF("当前配置：\n%+v", DefaultConfig)
		}
		return DefaultConfig
	}
	// 自定义配置（自定义配置覆盖于默认配置之上，所以自定义配置可以不写全，没写的会用默认配置）
	log.Info("正在加载自定义配置...")
	conf := DefaultConfig
	customConfigFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("读取自定义配置文件时发生错误：", err)
	}
	err = yaml.Unmarshal(customConfigFile, &conf)
	if err != nil {
		log.Fatal("反序列化默认配置文件时发生错误：", err)
	}
	log.Info("自定义配置加载完毕")
	out, err := yaml.Marshal(conf)
	log.InfoF("当前配置：\n%s", string(out))
	if err != nil {
		log.InfoF("当前配置：\n%+v", conf)
	}
	return conf
}
