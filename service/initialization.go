package service

import (
	"receiver_siem/hash"
	"receiver_siem/sender"
	"time"
)

type Initialization struct {
	Sender sender.Sender
	Key    string
}

func (init Initialization) Work() {
	init.Sender.Send(sender.InitInitializationMessage(init.Key, hash.ToMD5))
	time.Sleep(time.Second)
}
