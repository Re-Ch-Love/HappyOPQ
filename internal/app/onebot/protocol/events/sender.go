package events

type FriendMessageSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}

type GroupMessageSender struct {
	FriendMessageSender
	// TODO
}
