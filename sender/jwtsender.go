package sender

import (
	"bytes"
	"encoding/json"
	"net/http"
	"receiver_siem/entity/subject"
	"receiver_siem/hostinfo"
	"receiver_siem/token"
	"time"
)

type Message struct {
	Token string
	subject.Message
}

func (m Message) JSON() string {
	b, _ := json.Marshal(m)
	return string(b)
}

type JWTSender struct {
	HostServer string
	methods    map[string]CommandJWT
}

func InitJWTSender(hostSubject string) *JWTSender {
	return &JWTSender{HostServer: hostSubject,
		methods: map[string]CommandJWT{
			"init_receiver": CommandJWTPostForm{Address: hostSubject},
		}}
}

func (j *JWTSender) Send(message subject.Message) bool {
	resp, err := j.methods[message.TypeMessage].Command(Message{token.GetToken(), message}.JSON())
	if err != nil {
		return false
	}
	j.parse(resp)

	return resp.StatusCode == 200
}

func InitInitializationMessage() subject.Message {
	hostInfo := hostinfo.GetHostInfo()
	t := time.Now()
	return subject.Message{
		Message:     hostInfo.CodeName,
		TypeMessage: "init_receiver",
		HostName:    hostInfo.HostName,
		SystemOS:    hostInfo.HostOS,
		HostIP:      hostInfo.IPs,
		Time:        t,
	}
}

func (j *JWTSender) parse(resp *http.Response) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	str := buf.Bytes()
	m := make(map[string]interface{})
	json.Unmarshal(str, &m)
	_, ok := m["token"]
	if ok {
		token.SetToken(m["token"].(string))
	}
}
