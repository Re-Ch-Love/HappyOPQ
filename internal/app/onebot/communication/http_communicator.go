package communication

import (
	"HappyOPQ/pkg/log"
	"bytes"
	"io/ioutil"
	"net/http"
)

type HTTPCommunicator struct {
	URL string
}

func (c HTTPCommunicator) StartAPIServer() error {
	panic("implement me")
}

func (c HTTPCommunicator) Report(event []byte) error {
	resp, err := http.Post(c.URL,
		"application/x-www-form-urlencoded",
		bytes.NewReader(event))
	if err != nil {
		log.Error("向用户端转发事件时（发送请求时）出现错误：", err)
		return err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Error("向用户端转发事件时（关闭响应时）出现错误：", err)
		}
	}()
	// TODO 快速操作
	// 第一个返回值是读取出的内容
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("向用户端转发事件时（读取响应内容时）出现错误：", err)
		return err
	}
	return nil
}
