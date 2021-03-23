package opqbot

import (
	"HappyOPQ/internal/app/common"
	"os"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	eventCh := make(chan common.Convertible)
	flagCh := make(chan int)
	c := Connect("127.0.0.1", 8080, eventCh, flagCh)
	defer c.Close()
	go func() {
		time.Sleep(10 * time.Minute)
		os.Exit(0)
	}()
	<-flagCh
}
