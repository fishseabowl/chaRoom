package message

import (
	"fmt"
	"time"

	"github.com/fishseabowl/chatRoom/pkg/user"
)

// Stringer represents the wrap interface
type Stringer interface {
	String() string
}

// Message represents message structure in chat
type Message struct {
	date   time.Time  // timestamp
	text   string     // content of the message
	sender *user.User // user which sent messages
}

// String returns string representation of a message.
func (m *Message) String() string {
	return fmt.Sprintf("%s | %s - %s\r\n", m.date.Format("09/11/2020 20:01:01"), m.sender.Name(), m.text)
}

// NewMessage generates a message with information about the user, content and time.
func NewMessage(sender *user.User, content string, date time.Time) Stringer {
	return &Message{
		date:   date,
		text:   content,
		sender: sender,
	}
}
