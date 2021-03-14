package events

import (
	onebot "HappyOPQ/internal/app/onebot/protocol/events"
	"HappyOPQ/internal/pkg/richtext"
	"HappyOPQ/pkg/log"
	"encoding/json"
	"time"
)

type FriendMessageData struct {
	FromUin    int64       `json:"FromUin"`
	ToUin      int64       `json:"ToUin"`
	MsgType    string      `json:"MsgType"`
	MsgSeq     int         `json:"MsgSeq"`
	Content    string      `json:"Content"`
	RedBagInfo interface{} `json:"RedBaginfo"`
}

type FriendMessagePacket struct {
	WebConnID string            `json:"WebConnId"`
	Data      FriendMessageData `json:"Data"`
}

type FriendMessage struct {
	CurrentPacket FriendMessagePacket `json:"CurrentPacket"`
	CurrentQQ     int64               `json:"CurrentQQ"`
}

func (msg *FriendMessage) Bytes() []byte {
	cqMsg := richtext.OPQCode2CQCode(msg.CurrentPacket.Data.Content)
	var subType string
	if msg.CurrentPacket.Data.MsgType == "TextMsg" {
		subType = "friend"
	} else if msg.CurrentPacket.Data.MsgType == "TempSessionMsg" {
		subType = "group"
	} else {
		subType = "other"
	}
	msg2 := onebot.FriendMessage{
		Message: onebot.Message{
			Time:        time.Now().UnixNano(),
			SelfId:      msg.CurrentQQ,
			PostType:    "message",
			MessageType: "private",
			SubType:     subType,
			MessageId:   int32(msg.CurrentPacket.Data.MsgSeq),
			UserId:      msg.CurrentPacket.Data.FromUin,
			Message:     richtext.RemoveCQCode(cqMsg),
			RawMessage:  cqMsg,
			Font:        0,
			Sender: onebot.FriendMessageSender{
				MessageSender: onebot.MessageSender{
					UserId:   msg.CurrentPacket.Data.FromUin,
					Nickname: "",
					Sex:      "unknown",
					Age:      0,
				},
			},
		},
	}
	bytes, err := json.Marshal(msg2)
	if err != nil {
		log.Error("转换消息格式时发生错误：", err)
		return nil
	}
	return bytes
}
