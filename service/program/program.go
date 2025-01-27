package program

import (
	"receiver_siem/config"
	"receiver_siem/hostinfo"
	"receiver_siem/service/analysis"
	service "receiver_siem/service/api"
	"receiver_siem/service/init"
	"time"
)

type Program struct {
	InitializationService init.Initialization
	ApiService            service.ApiService
	AnalysisService       analysis.Analysis
}

func InitProgram(fileName string) *Program {
	conf, err := config.ReadFromFile(fileName)
	if err != nil {
		panic(err)
	}
	hostinfo.HostInfoInit(conf.TelegramBot.CodeName)
	return &Program{}
}

func (p *Program) Work() {
	p.InitializationService.Work()
	time.Sleep(time.Second)
	go p.ApiService.Work()
	p.AnalysisService.Work()
}
