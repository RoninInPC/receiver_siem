package init

import (
	"receiver_siem/sender"
	"time"
)

type Initialization struct {
	Sender sender.Sender
}

func (init Initialization) Work() {
	init.Sender.Send(sender.InitInitializationMessage())
	time.Sleep(time.Second)
}
