package communication

import (
	"HappyOPQ/internal/app/common"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	eventCh := make(chan common.Convertible)
	c := Connect("127.0.0.1", 8080, eventCh)
	defer c.Close()
	time.Sleep(10 * time.Minute)
}
