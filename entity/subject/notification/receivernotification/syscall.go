package receivernotification

import (
	"receiver_siem/entity/subject"
)

type Syscall struct {
	subject.Syscall
	BaseNotification
}

func (s Syscall) GetInfo() string {
	//TODO implement me
	panic("implement me")
}

func (s Syscall) GetInfoMarkdown() string {
	//TODO implement me
	panic("implement me")
}

func (s Syscall) GetProcessPID() string {
	return s.PID
}

func (s Syscall) GetUsername() string {
	return s.Username
}
