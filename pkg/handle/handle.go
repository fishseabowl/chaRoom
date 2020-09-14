package handler

import (
	"bytes"
	"fmt"
	"io"

	"github.com/fishseabowl/chatRoom/pkg/usrctx"
)

// Handler is a wrapper for Serve
type Handler interface {
	Serve(ctx usrctx.Context)
}

const (
	defaultExitCommandName = "/quit"
	lineEnd                = "\r\n"
)

// ChatHandler is a base handler for serving a chat.
type ChatHandler struct {
	ExitCommandName string
}

// NewChatHandler creates handler which can be used on server.
func NewChatHandler() Handler {
	return &ChatHandler{
		ExitCommandName: defaultExitCommandName,
	}
}

// Serve is method which processes all requests from user.
func (h *ChatHandler) Serve(ctx usrctx.Context) {
	buf := make([]byte, 50)
	cmdLine := new(bytes.Buffer)

	usrCtx, ok := ctx.(*usrctx.UserContext)
	if !ok {
		panic("Provided context is wrong. Should be UserContext")
	}

	chatInst := usrCtx.Chat()
	usr := usrCtx.User()
	chann := usrCtx.Channel()

	// Send notification about new user
	msg := chann.AddMessage(usr, fmt.Sprintf("User %s connected to channel %s", usr.Name(), chann.Name()))
	chann.SendMessage(usr, msg)

	// Say hello to new user
	io.WriteString(usr, chatInst.WelcomeMessage())
	io.WriteString(usr, fmt.Sprintf("Nick %s, welcome in a channel %s. There is %d other users\r\n", usr.Name(), chann.Name(), len(chann.Users())-1))

	// Main handler loop
	for {
		n, err := usr.Read(buf)
		if n <= 0 && err == nil {
			// Nothing to do, let's check again
			continue
		} else if n <= 0 && err != nil {
			// Some error happended and not connection closed. Log it and end.
			if err != io.EOF {
				panic(err)
			}
			break
		}
		buf = bytes.Trim(buf, " ")
		h.writeBuf(cmdLine, buf)

	}
}

func (h ChatHandler) isEmptyLine(line *bytes.Buffer) bool {
	lineEndOnly := bytes.Equal(line.Bytes()[:2], []byte(lineEnd))

	if line.Len() > 2 && lineEndOnly {
		return true
	}

	return false
}

func (h ChatHandler) clear(line *bytes.Buffer) {
	line.Reset()
}

// isCompleteLine checks if there is \n in text send.
// Checks that on buffer to speed up processing and don't go over growing line.
func (h ChatHandler) isCompleteLine(buf []byte) bool {
	if bytes.Index(buf, []byte(lineEnd)) > -1 {
		return true
	}

	return false
}

func (h ChatHandler) writeBuf(line *bytes.Buffer, buf []byte) {
	writeTo := bytes.Index(buf, []byte(lineEnd))

	if writeTo > 0 {
		line.Write(buf[:writeTo])
	} else {
		line.Write(buf)
	}
}
