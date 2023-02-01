package main

import (
	"time"

	"github.com/humbertovnvarro/vrchatgo/pkg/vrcosc"
	"github.com/hypebeast/go-osc/osc"
)

func main() {
	client := vrcosc.New(&vrcosc.ClientOptions{
		Hostname:     "localhost",
		SenderPort:   9000,
		ReceiverPort: 9001,
	})
	for {
		time.Sleep(time.Second)
		client.Sender.Send(osc.NewMessage(string(vrcosc.Jump), 1))
	}
}
