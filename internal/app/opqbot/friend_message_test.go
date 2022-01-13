package opqbot

import (
	"testing"
)

func TestFriendMessage_Report(t *testing.T) {
	_ = FriendMessage{
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
	//logger.Debug(msg.Convert())
}
