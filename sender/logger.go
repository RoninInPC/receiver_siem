package sender

import (
	"log"
	"receiver_siem/entity/subject"
)

type Logger struct {
}

func (logger *Logger) Send(subject subject.Message) bool {
	log.Println(subject)
	return true
}
