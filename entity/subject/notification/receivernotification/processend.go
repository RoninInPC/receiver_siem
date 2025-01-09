package receivernotification

import (
	"encoding/json"
	"fmt"

	"receiver_siem/entity/subject"
	"receiver_siem/hash"
)

type ProcessEnd struct {
	Process subject.Process
	BaseNotification
}

func (p ProcessEnd) GetInfo() string {
	//TODO implement me
	panic("implement me")
}

func (p ProcessEnd) GetInfoMarkdown() string {
	//TODO implement me
	panic("implement me")
}

func (p ProcessEnd) JSON() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (p ProcessEnd) Type() subject.SubjectType {
	return ProcessEndT
}

func (p ProcessEnd) Name() string {
	return fmt.Sprintf("Процесс %s завершился.",
		p.Process.Name())
}

func (p ProcessEnd) Hash(hash hash.Hash) string {
	return hash(p.JSON())
}
