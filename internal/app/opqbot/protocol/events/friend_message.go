package events

import (
	onebot "HappyOPQ/internal/app/onebot/protocol/events"
	"HappyOPQ/internal/app/richtext"
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

func (msg *FriendMessage) Convert() (int64, interface{}) {
	cqMsg := richtext.OPQCode2CQCode(msg.CurrentPacket.Data.Content)
	var subType string
	if msg.CurrentPacket.Data.MsgType == "TextMsg" {
		subType = "friend"
	} else if msg.CurrentPacket.Data.MsgType == "TempSessionMsg" {
		subType = "group"
	} else {
		subType = "other"
	}
	return msg.CurrentQQ, onebot.FriendMessage{
		Time:        time.Now().Unix(),
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
			UserId:   msg.CurrentPacket.Data.FromUin,
			Nickname: "",
			Sex:      "unknown",
			Age:      0,
		},
	}
}
