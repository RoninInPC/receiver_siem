package command

import (
	tm "github.com/and3rson/telemux/v2"
	"github.com/gin-gonic/gin"
)

type ApiAction interface {
	Action(*gin.Context)
}

type TelegramAction interface {
	Action(*tm.Update)
}
