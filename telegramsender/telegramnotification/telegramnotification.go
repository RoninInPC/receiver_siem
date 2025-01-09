package telegramnotification

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramSimple struct {
	api *tgbotapi.BotAPI
}

func Init(api *tgbotapi.BotAPI) *TelegramSimple {
	return &TelegramSimple{api: api}
}

func (t TelegramSimple) Send(id int64, message string) bool {
	msg := tgbotapi.NewMessage(id, message)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := t.api.Send(msg)
	return err == nil
}
