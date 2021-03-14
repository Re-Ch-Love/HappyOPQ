package events

type MessageSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}

type FriendMessageSender struct {
	MessageSender
}

type GroupMessageSender struct {
	MessageSender
	// TODO
}
