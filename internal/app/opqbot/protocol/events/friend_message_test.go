package events

import (
	"HappyOPQ/pkg/log"
	"testing"
)

func TestFriendMessage_Report(t *testing.T) {
	msg := FriendMessage{
		CurrentPacket: FriendMessagePacket{
			WebConnID: "",
			Data: FriendMessageData{
				FromUin:    0,
				ToUin:      0,
				MsgType:    "",
				MsgSeq:     0,
				Content:    "",
				RedBagInfo: nil,
			},
		},
		CurrentQQ: 0,
	}
	log.Debug(string(msg.Bytes()))
}
