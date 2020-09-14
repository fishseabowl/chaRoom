package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fishseabowl/chatRoom/pkg/chat"
	handler "github.com/fishseabowl/chatRoom/pkg/handle"
	"github.com/fishseabowl/chatRoom/pkg/server"
	"github.com/fishseabowl/chatRoom/pkg/usrctx"
)

//OnServer is a telent server
func OnServer(addr string) error {

	chat := chat.NewChat()
	ctx := usrctx.NewContext()
	ctx.SetAttribute("chat", chat)

	srv := server.NewServer(addr, handler.NewChatHandler(), ctx)
	if err := srv.ListenAndServe(); nil != err {
		panic(err)
	}

	return nil
}

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go OnServer("localhost:telnet")
	<-sigs
	fmt.Println("existing")
}
