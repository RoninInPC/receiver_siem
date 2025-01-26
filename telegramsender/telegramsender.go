package telegramsender

type TelegramSender[A any] interface {
	Send(id int64, a A) bool
	SendSeveral(ids int64, a []A) bool
}
