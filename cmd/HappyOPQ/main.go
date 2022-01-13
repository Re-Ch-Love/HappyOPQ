package main

import (
	"HappyOPQ/internal/app/config"
	onebotComm "HappyOPQ/internal/app/onebot/communication"
	"HappyOPQ/internal/app/opqbot"
	"HappyOPQ/pkg/utils"
	"flag"
	"os"
	"os/exec"
	"time"
)

var logger = utils.NewDefaultLogger()

func init() {
	logger.Tag = "main"
}

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
	// 启动OPQBot
	cmd := exec.Command(conf.OPQBot.ExecPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	defer func() {
		err := cmd.Wait()
		if err != nil {
			logger.PanicErr(err)
		}
	}()
	go func() {
		err := cmd.Start()
		if err != nil {
			logger.PanicErr(err)
		}
	}()
	// 与OPQBot建立连接
	client := opqbot.NewCommunicator(&conf.OPQBot)
	time.Sleep(time.Second * 10)
	finishSignalChan := make(chan struct{})
	go client.Run(finishSignalChan)
	defer client.Close()
	<-finishSignalChan
	// 与用户端建立连接
	c := conf.OneBot
	if c.HTTP.Enabled {
		logger.Info("使用 HTTP 与 OneBot 通信")
		communicator := onebotComm.HTTPCommunicator{URL: c.HTTP.URL}
		// 阻塞
		for {
			e := <-client.EventChan
			_ = communicator.Report(e.Convert())
		}
	}
	if c.PositiveWebSocket.Enabled {
		// TODO
		logger.Info("抱歉，HappyOPQ 暂不支持使用正向 WebSocket 与 OneBot 通信")
	}
	if c.ReverseWebSocket.Enabled {
		// TODO
		logger.Info("抱歉，HappyOPQ 暂不支持使用反向 WebSocket 与 OneBot 通信")
	}
}
