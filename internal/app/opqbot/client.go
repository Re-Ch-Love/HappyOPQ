package opqbot

import (
	"HappyOPQ/internal/app/common"
	"HappyOPQ/internal/app/common/retry"
	"HappyOPQ/internal/app/config"
	"HappyOPQ/pkg/utils"
	sio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

var logger = utils.NewDefaultLogger()

func init() {
	logger.Tag = "client"
}

const (
	OnFriendMsgs = "OnFriendMsgs"
	OnGroupMsgs  = "OnGroupMsgs"
	OnEvents     = "OnEvents"
)

type Communicator struct {
	Host      string
	Port      int
	FlagChan  chan int
	EventChan chan common.Convertible
	client    *sio.Client
	retry     retry.Strategy
}

func (c *Communicator) Close() {
	if c.client == nil {
		return
	}
	c.client.Close()
}

// NewCommunicator
// TODO 配置中加入对各种错误的重试策略
func NewCommunicator(conf *config.OPQBotConfig) *Communicator {
	c := &Communicator{
		Host:      conf.Host,
		Port:      conf.Port,
		FlagChan:  make(chan int),
		EventChan: make(chan common.Convertible),
		client:    nil,
		retry:     retry.Timed(3, 3), //todo
	}
	return c
}

func (c *Communicator) connect() error {
	client, err := sio.Dial(sio.GetUrl(c.Host, c.Port, false), transport.GetDefaultWebsocketTransport())
	if err != nil {
		return err
	}
	c.client = client
	return nil
}
func (c *Communicator) init() {
	utils.Must(c.client.On(sio.OnConnection, func(_c *sio.Channel) {
		logger.Info("成功与 OPQBot 连接，ID 为", c.client.Id())
		c.FlagChan <- ConnectionSucceed
	}))

	utils.Must(c.client.On(sio.OnDisconnection, func(_c *sio.Channel) {
		logger.Error("与 OPQBot 断开连接")
		c.FlagChan <- ConnectionTerminate
	}))

	utils.Must(c.client.On(sio.OnError, func(_c *sio.Channel) {
		logger.Error("与 OPQBot 的连接发生错误")
		c.FlagChan <- ConnectionTerminate
	}))

	utils.Must(c.client.On(OnFriendMsgs, func(conn *sio.Channel, msg FriendMessage) {
		logger.Infof("收到 OnFriendMsgs 事件：%+v", msg)
		c.EventChan <- &msg
	}))

	utils.Must(c.client.On(OnGroupMsgs, func(c *sio.Channel, msg interface{}) {
		logger.Infof("收到 OnGroupMsgs 事件：%+v", msg)
	}))

	utils.Must(c.client.On(OnEvents, func(conn *sio.Channel, msg interface{}) {
		logger.Infof("收到 OnEvents 事件：%+v", msg)
	}))
}

func (c *Communicator) Run(finishSignal chan struct{}) {
	defer func() {
		if err := recover(); err != nil {
			logger.PanicErr(err)
		}
	}()
	utils.Must(c.connect())
	c.init()
	finishSignal <- struct{}{}
	close(finishSignal)
	c.handleFlags()
}

func (c *Communicator) handleFlags() {
	for {
		f := <-c.FlagChan
		switch f {
		case ConnectionTerminate:
			utils.Must(c.retry.On(c.connect))
			c.init()
		}
	}
}
