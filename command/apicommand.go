package command

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"receiver_siem/entity/subject"
	"receiver_siem/entity/subject/notification/receivernotification"
	"receiver_siem/token"
	"strings"
)

type ApiAction interface {
	Action(*gin.Context)
}

type CommandNotification struct {
	CheckerHostName string
	Channel         chan receivernotification.Notification
}

func (action CommandNotification) Action(g *gin.Context) {
	request := g.Request
	body, _ := io.ReadAll(request.Body)
	defer request.Body.Close()
	log.Println(string(body))
	jsoned := string(body)
	jsoned = strings.Replace(jsoned, "%7B", "{", -1)
	jsoned = strings.Replace(jsoned, "%22", "\"", -1)
	jsoned = strings.Replace(jsoned, "%3A", ":", -1)
	jsoned = strings.Replace(jsoned, "%2C", ",", -1)
	jsoned = strings.Replace(jsoned, "%7D", "}", -1)
	jsoned = strings.Replace(jsoned, "json=", "", -1)
	m := subject.Message{}
	json.Unmarshal([]byte(jsoned), &m)
	if m.HostName == action.CheckerHostName && m.Token == token.GetToken() {
		action.Channel <- receivernotification.JsonedToNotification(m.Json, m.TypeSubject)
	}
}
