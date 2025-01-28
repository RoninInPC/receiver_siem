package program

import (
	"receiver_siem/command"
	"receiver_siem/config"
	"receiver_siem/entity/subject/notification/receivernotification"
	user2 "receiver_siem/entity/user"
	"receiver_siem/hostinfo"
	"receiver_siem/sender"
	"receiver_siem/service/analysis"
	service "receiver_siem/service/api"
	"receiver_siem/service/initialization"
	"receiver_siem/service/telegram"
	"receiver_siem/storageuser/redisusers"
	"receiver_siem/telegramsender/telegramnotification"
	"time"
)

type Program struct {
	InitializationService initialization.Initialization
	ApiService            service.ApiService
	TelegramBot           telegram.TelegramBot
	AnalysisService       analysis.Analysis
}

func InitProgram(fileName string) *Program {
	conf, err := config.ReadFromFile(fileName)
	if err != nil {
		panic(err)
	}
	hostinfo.HostInfoInit(conf.TelegramBot.Port)
	initService := initialization.Initialization{Sender: sender.InitJWTSender(conf.Host.HostAddressServer)}
	channel := make(chan receivernotification.Notification)
	apiService := service.InitApiService(
		conf.TelegramBot.Address+":"+conf.TelegramBot.Port,
		conf.Host.HostCheck,
		channel)
	usersRedis := redisusers.InitRedisUserTable(
		conf.RedisUsers.Address,
		conf.RedisUsers.Password,
		conf.RedisUsers.DB)
	usersRedis.Append(*user2.InitUser(
		conf.AdminUser.UserName,
		conf.AdminUser.FirstName,
		conf.AdminUser.SecondName,
		conf.AdminUser.ID,
		conf.AdminUser.ChatID,
		conf.AdminUser.Rights))
	telegramBot, err := telegram.InitBot(conf.TelegramBot.Token)
	if err != nil {
		panic(err)
	}
	senderTelegram := telegramnotification.Init(telegramBot.BotApi)
	telegramBot.AddCommand(
		telegram.MakeCommandByFilterDefault(
			"start",
			"начнём?",
			command.StartTelegram{
				Sender:          senderTelegram,
				StorageUser:     usersRedis,
				StartPassString: conf.TelegramBot.AccessStart}),
	)
	return &Program{initService, apiService, *telegramBot, analysis.Init(
		channel,
		usersRedis,
		senderTelegram,
		time.Minute*5)}
}

func (p *Program) Work() {
	p.InitializationService.Work()
	go p.ApiService.Work()
	go p.AnalysisService.Work()
	p.TelegramBot.Work()
}
