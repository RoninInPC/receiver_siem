package receivernotification

import (
	"encoding/json"
	"fmt"
	"receiver_siem/entity/subject"
	"receiver_siem/hash"
)

type UserNew struct {
	User subject.User
	BaseNotification
}

func (u UserNew) GetInfo() string {
	//TODO implement me
	panic("implement me")
}

func (u UserNew) GetInfoMarkdown() string {
	//TODO implement me
	panic("implement me")
}

func (u UserNew) JSON() string {
	bytes, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (u UserNew) Type() subject.SubjectType {
	return UserNewT
}

func (u UserNew) Name() string {
	return fmt.Sprintf("Пользователь %s создан %s(%s) в процессе %s (%s).",
		u.User.Username,
		u.Who.Username, u.Who.Uid,
		u.WhoProcess.PID, u.WhoProcess.NameProcess)
}

func (u UserNew) Hash(hash hash.Hash) string {
	return hash(u.JSON())
}
