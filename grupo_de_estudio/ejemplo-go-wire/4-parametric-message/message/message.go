package message

// Message is what greeters will use to greet guests.
type Message string

// NewMessage creates a default Message.
func NewMessage(phrase string) Message {
	return Message(phrase)
}
