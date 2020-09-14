package chat

import (
	"sync"

	"github.com/fishseabowl/chatRoom/pkg/channel"
)

// Chat represtns tje information about all the chat room
type Chat struct {
	channels       map[string]*channel.Channel
	welcomeMessage string
	exitMessage    string
	mutex          sync.RWMutex
}

// NewChat creates new instance of a chat.
func NewChat() *Chat {
	chat := &Chat{
		channels:       make(map[string]*channel.Channel),
		welcomeMessage: "\r\nHello!\r\n",
		exitMessage:    "\r\nGoodbye!\r\n",
	}

	chann := channel.NewChannel(channel.DefaultChannelName)
	chat.AddChannel(chann)

	return chat
}

// WelcomeMessage returns messages when user connect to the caht room.
func (c *Chat) WelcomeMessage() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.welcomeMessage
}

// ExitMessage returns messages when user leave the caht room.
func (c *Chat) ExitMessage() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.exitMessage
}

// AddChannel registers new chat room on the server.
func (c *Chat) AddChannel(chann *channel.Channel) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.channels[chann.Name()] = chann
}

// Channels returns a map of all chat room registered on the server.
func (c *Chat) Channels() map[string]*channel.Channel {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.channels
}
