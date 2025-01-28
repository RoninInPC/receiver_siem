package command

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/url"
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
	m := contextToMessage(g)
	if m.HostName == action.CheckerHostName && m.Token == token.GetToken() {
		action.Channel <- receivernotification.JsonedToNotification(m.Json, m.TypeSubject)
	}
}

func contextToMessage(g *gin.Context) subject.Message {
	request := g.Request
	body, _ := io.ReadAll(request.Body)
	defer request.Body.Close()
	log.Println(string(body))
	jsoned := string(body)
	jsoned, _ = url.QueryUnescape(jsoned)
	//jsoned = strings.Replace(jsoned, "\\\"", "\"", -1)
	jsoned = strings.Replace(jsoned, "json=", "", -1)
	m := subject.Message{}
	err := json.Unmarshal([]byte(jsoned), &m)
	if err != nil {
		log.Println(err.Error())
	}
	return m
}
