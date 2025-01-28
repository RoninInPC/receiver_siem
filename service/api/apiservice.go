package service

import (
	"receiver_siem/api"
	"receiver_siem/command"
	"receiver_siem/entity/subject/notification/receivernotification"
)

type Method int

var (
	POST    Method = 0
	GET     Method = 1
	HEAD    Method = 2
	OPTIONS Method = 3
	PATCH   Method = 4
	PUT     Method = 5
	DELETE  Method = 6
)

type PathWork struct {
	Method Method
	Path   string
	Action command.ApiAction
}

type ApiService struct {
	Address  string
	API      api.Api
	Commands []PathWork
}

func InitApiService(address, checkHostname string,
	channel chan receivernotification.Notification) ApiService {
	return ApiService{API: api.InitApi(), Address: address, Commands: []PathWork{
		{POST, "/api/alert", command.CommandNotification{CheckerHostName: checkHostname, Channel: channel}},
	}}
}

func (a ApiService) Work() {
	if &a != nil {
		a.API = api.InitApi()
	}
	for _, c := range a.Commands {
		switch c.Method {
		case POST:
			a.API.Post(c.Path, c.Action)
		case GET:
			a.API.Get(c.Path, c.Action)
		case HEAD:
			a.API.Head(c.Path, c.Action)
		case OPTIONS:
			a.API.Options(c.Path, c.Action)
		case PATCH:
			a.API.Patch(c.Path, c.Action)
		case PUT:
			a.API.Patch(c.Path, c.Action)
		case DELETE:
			a.API.Delete(c.Path, c.Action)
		default:
			continue
		}
	}
	a.API.Run(a.Address)
}
