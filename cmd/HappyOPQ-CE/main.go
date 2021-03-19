package main

import (
	"HappyOPQ/internal/app/common"
	"HappyOPQ/internal/app/config"
	"HappyOPQ/internal/app/onebot/communication"
	OPQBotCommunicators "HappyOPQ/internal/app/opqbot/communication"
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
	eventCh := make(chan common.Convertible)
	flagCh := make(chan int)
	// 加载配置
	conf := config.LoadConfig(*configPath)
	// 与OPQBot建立连接
	client := OPQBotCommunicators.Connect(conf.OPQBot.Host, conf.OPQBot.Port, eventCh, flagCh)
	defer client.Close()
	// 与用户端建立连接.
	c := conf.OneBot
	if c.HTTP.Enabled {
		log.Info("使用 HTTP 与 OneBot 通信")
		communicator := communication.HTTPCommunicator{URL: c.HTTP.URL}
		go func() {
			for {
				e := <-eventCh
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
	// TODO
	/*var reconnectionCount int
	for {
		select {
		case f := <-flagCh:
			switch f {
			// 处理与 OPQBot 连接的的客户端的异常，即重连
			case OPQBotCommunicators.ConnectionTerminate:
				time.Sleep(10)
				reconnectionCount++
				if reconnectionCount <= constants.MaxReconnectionTimes {
					log.InfoF("将在 %d 秒后开始尝试重连", constants.ReconnectionInterval)
					time.Sleep(time.Second * constants.ReconnectionInterval)
					log.InfoF("重连中...第 %d/%d 次尝试", reconnectionCount, constants.MaxReconnectionTimes)
					client = OPQBotCommunicators.Connect(conf.OPQBot.Host, conf.OPQBot.Port, eventCh, flagCh)
					log.Info("重连成功！")
				} else {
					log.Fatal("与 OPQBot 重连失败，退出程序！")
				}
			case OPQBotCommunicators.ConnectionSucceed:
				reconnectionCount = 0
			}
		}
	}*/
}
