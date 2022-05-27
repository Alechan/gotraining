package message

// Message is what greeters will use to greet guests.
type Message string

// NewMessage creates a default Message.
func NewMessage() Message {
	return Message("Hi there!")
}
