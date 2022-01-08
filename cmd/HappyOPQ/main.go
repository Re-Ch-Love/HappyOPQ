package main

import (
	"HappyOPQ/internal/app/config"
	"HappyOPQ/internal/app/onebot/communication"
	"HappyOPQ/internal/app/opqbot"
	"HappyOPQ/pkg/log"
	"flag"
	"os"
)

var configPath = flag.String("c", "./opq-onebot.yml", "配置文件的路径")

func init() {
	// 当参数为`help`或`?`时输出帮助信息
	if len(os.Args) > 1 {
		if os.Args[1] == "help" || os.Args[1] == "?" {
			flag.Usage()
		}
	}
	// 解析命令行参数
	flag.Parse()
}

func main() {
	// 加载配置
	conf := config.LoadConfig(*configPath)
	// 与OPQBot建立连接

	client := opqbot.NewCommunicator(&conf.OPQBot)
	client.Init()
	defer client.Close()
	// 与用户端建立连接.
	c := conf.OneBot
	if c.HTTP.Enabled {
		log.Info("使用 HTTP 与 OneBot 通信")
		communicator := communication.HTTPCommunicator{URL: c.HTTP.URL}
		go func() {
			for {
				e := <-client.EventChan
				_ = communicator.Report(e.Convert())
			}
		}()
	}
	if c.PositiveWebSocket.Enabled {
		// TODO
		log.Info("抱歉，HappyOPQ 暂不支持使用正向 WebSocket 与 OneBot 通信")
	}
	if c.ReverseWebSocket.Enabled {
		// TODO
		log.Info("抱歉，HappyOPQ 暂不支持使用反向 WebSocket 与 OneBot 通信")
	}
	// 阻塞
	<-make(chan interface{})
}
