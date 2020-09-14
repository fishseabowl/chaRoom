package channel

import (
	"io"
	"net"
	"sync"
	"time"

	"github.com/fishseabowl/chatRoom/pkg/message"
	"github.com/fishseabowl/chatRoom/pkg/user"
)

// DefaultChannelName represents the default channel name
const (
	DefaultChannelName string = "default"
)

// Call represents a function Call.
type Call func(*user.User)

// Channel represents base functionality of a server.
type Channel struct {
	name     string                  // Name of the chat room
	users    map[net.Conn]*user.User // Map of all users' connection
	messages []message.Stringer      // Slice with messages sent to the chat room
	mutex    sync.RWMutex            // mutex for access to messages and users
}

// NewChannel creates instance of new chat room
func NewChannel(name string) *Channel {
	users := make(map[net.Conn]*user.User)
	messages := make([]message.Stringer, 5, 5)
	return &Channel{
		name:     name,
		users:    users,
		messages: messages,
	}
}

// Name represnts the room Name
func (c *Channel) Name() string {
	return c.name
}

// AddUser registers a new user into chat room.
func (c *Channel) AddUser(user *user.User) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.users[user.Conn()] = user
}

// DelUser drop a  user from chat room.
func (c *Channel) DelUser(usr *user.User) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.users, usr.Conn())
}

// Users return users list in chat room.
func (c *Channel) Users() map[net.Conn]*user.User {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.users
}

// Call iterates over all registered users in chat room.
func (c *Channel) Call(fn Call) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	for _, user := range c.users {
		fn(user)
	}
}

// AddMessage registers new message in chat.
func (c *Channel) AddMessage(sender *user.User, content string) message.Stringer {
	msg := message.NewMessage(sender, content, time.Now())
	c.addStringer(msg)

	return msg
}

func (c *Channel) addStringer(msg message.Stringer) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.messages = append(c.messages, msg)
}

// SendMessage sends stringer to all the users in the chat room.
func (c *Channel) SendMessage(sender *user.User, msg message.Stringer) {
	c.Call(func(u *user.User) {
		if u != sender {
			io.WriteString(u, msg.String())
		}
	})
}
