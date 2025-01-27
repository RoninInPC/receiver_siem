package command

import (
	"fmt"
	tm "github.com/and3rson/telemux/v2"
	"receiver_siem/command/updatetelegram"
	user2 "receiver_siem/entity/user"
	"receiver_siem/other"
	"receiver_siem/storageuser"
	"receiver_siem/telegramsender"
	"strconv"
	"strings"
)

const (
	WrongStart = "Вы не допущены до бота."
	AccessUser = "Пользователь %s впервые воспользовался ботом."
)

type TelegramAction interface {
	Action(*tm.Update)
}

type StartTelegram struct {
	Sender          telegramsender.TelegramSender[string]
	StorageUser     storageuser.StorageUser
	StartPassString string
}

func (s StartTelegram) Action(u *tm.Update) {
	message := updatetelegram.GetMessage(u)
	text := strings.Split(message.Text, " ")
	if len(text) < 2 {
		s.Sender.Send(updatetelegram.GetUserId(u), WrongStart)
	}
	if text[1] != s.StartPassString {
		s.Sender.Send(updatetelegram.GetUserId(u), WrongStart)
	}
	idString := strconv.FormatInt(message.From.ID, 10)
	if !s.StorageUser.Contains(idString) {
		user := user2.InitUser(
			message.From.UserName,
			message.From.FirstName,
			message.From.LastName,
			other.GetUserId(u),
			other.GetUserId(u),
			5)
		s.StorageUser.Append(
			*user,
		)
		for _, user := range s.StorageUser.GetUserByRights(10) {
			s.Sender.Send(
				user.Info.ChatId,
				fmt.Sprintf(
					AccessUser,
					user.ToString(true),
				),
			)
		}
	}
}
