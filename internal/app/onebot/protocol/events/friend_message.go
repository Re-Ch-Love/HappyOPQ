package events

type FriendMessage struct {
	Time        int64               `json:"time"`
	SelfId      int64               `json:"self_id"`
	PostType    string              `json:"post_type"`
	MessageType string              `json:"message_type"`
	SubType     string              `json:"sub_type"`
	MessageId   int32               `json:"message_id"`
	UserId      int64               `json:"user_id"`
	Message     string              `json:"message"`
	RawMessage  string              `json:"raw_message"`
	Font        int32               `json:"font"`
	Sender      FriendMessageSender `json:"sender"`
}
