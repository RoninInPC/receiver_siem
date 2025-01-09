package receivernotification

import (
	"encoding/json"
	"fmt"

	"receiver_siem/entity/subject"
	"receiver_siem/hash"
)

type ProcessNew struct {
	Process subject.Process
	BaseNotification
}

func (p ProcessNew) GetInfo() string {
	//TODO implement me
	panic("implement me")
}

func (p ProcessNew) GetInfoMarkdown() string {
	//TODO implement me
	panic("implement me")
}

func (p ProcessNew) JSON() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (p ProcessNew) Type() subject.SubjectType {
	return ProcessNewT
}

func (p ProcessNew) Name() string {
	return fmt.Sprintf("Процесс %s создан %s(%s) в процессе %s (%s).",
		p.Process.Name(),
		p.Who.Username, p.Who.Uid,
		p.WhoProcess.PID, p.WhoProcess.NameProcess)
}

func (p ProcessNew) Hash(hash hash.Hash) string {
	return hash(p.JSON())
}
