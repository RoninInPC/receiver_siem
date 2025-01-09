package telegramsender

type TelegramSender[A any] interface {
	Send(id int64, a A) bool
}
