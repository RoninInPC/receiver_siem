package command

import tm "github.com/and3rson/telemux/v2"

type TelegramAction interface {
	Action(*tm.Update)
}
