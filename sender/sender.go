package sender

import (
	"receiver_siem/entity/subject"
)

type Sender interface {
	Send(message subject.Message) bool
}
