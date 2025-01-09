package command

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"receiver_siem/entity/subject"
	"receiver_siem/entity/subject/notification/receivernotification"
	"receiver_siem/token"
)

type ApiAction interface {
	Action(*gin.Context)
}

type CommandNotification struct {
	CheckerHostName string
	Channel         chan receivernotification.Notification
}

func (action CommandNotification) Action(g *gin.Context) {
	j := g.Param("json")
	m := subject.Message{}
	json.Unmarshal([]byte(j), &m)
	if m.HostName == action.CheckerHostName && m.Token == token.GetToken() {
		action.Channel <- receivernotification.JsonedToNotification(m.Json, m.TypeSubject)
	}
}
