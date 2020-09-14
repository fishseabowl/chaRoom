package usrctx

import (
	"github.com/fishseabowl/chatRoom/pkg/channel"
	"github.com/fishseabowl/chatRoom/pkg/chat"
	"github.com/fishseabowl/chatRoom/pkg/user"
)

// NewUserContext creates a context with access to special fields based on the
// attributes.
func NewUserContext() *UserContext {
	return &UserContext{
		NewContext(),
	}
}

// UserContext is a special context wrapper.
type UserContext struct {
	Context
}

// Chat provides access to attribute.
func (ctx UserContext) Chat() *chat.Chat {
	attr, err := ctx.Attribute("chat")
	if err != nil {
		panic(err)
	}

	chatInst, ok := attr.(*chat.Chat)
	if ok == false {
		panic("Chat instance is not available")
	}

	return chatInst
}

// User provides access to attribute.
func (ctx *UserContext) User() *user.User {
	attr, err := ctx.Attribute("user")

	if err != nil {
		panic(err)
	}

	usr, ok := attr.(*user.User)
	if ok == false {
		panic("User instance is not available")
	}

	return usr
}

// Channel provides access to attribute.
func (ctx *UserContext) Channel() *channel.Channel {
	attr, err := ctx.Attribute("channel")

	if err != nil {
		panic(err)
	}

	ch, ok := attr.(*channel.Channel)
	if ok == false {
		panic("User's channel is not available")
	}

	return ch
}
