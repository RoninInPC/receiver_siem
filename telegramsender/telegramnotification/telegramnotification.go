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

func (t TelegramSimple) SendSeveral(id int64, messages []string) bool {
	for _, message := range messages {
		if !t.Send(id, message) {
			return false
		}
	}
	return true
}
