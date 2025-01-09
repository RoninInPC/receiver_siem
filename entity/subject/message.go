package subject

import (
	"encoding/json"
	"receiver_siem/hash"
	"receiver_siem/hostinfo"
	"time"
)

type Message struct {
	Token       string
	Message     string
	TypeMessage string
	HostName    string
	SystemOS    string
	HostIP      []string
	Time        time.Time
	TypeSubject SubjectType
	Json        string
}

func InitMessage(
	token string,
	message string,
	typeMessage string,
	hostInfo hostinfo.HostInfo,
	subject Subject) Message {
	return Message{
		token,
		message,
		typeMessage,
		hostInfo.HostName,
		hostInfo.HostOS,
		hostInfo.IPs,
		time.Now(),
		subject.Type(),
		subject.JSON()}
}

func (m Message) JSON() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (m Message) Type() SubjectType {
	return MessageT
}

func (m Message) Name() string {
	return m.JSON()
}

func (m Message) Hash(hash hash.Hash) string {
	return hash(m.JSON())
}
