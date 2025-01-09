package receivernotification

import (
	"encoding/json"
	"fmt"

	"receiver_siem/entity/subject"
	"receiver_siem/hash"
)

type ProcessUpdate struct {
	ProcessBefore subject.Process
	ProcessAfter  subject.Process
	BaseNotification
}

func (p ProcessUpdate) GetInfo() string {
	//TODO implement me
	panic("implement me")
}

func (p ProcessUpdate) GetInfoMarkdown() string {
	//TODO implement me
	panic("implement me")
}

func (p ProcessUpdate) JSON() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (p ProcessUpdate) Type() subject.SubjectType {
	return ProcessUpdateT
}

func (p ProcessUpdate) Name() string {
	return fmt.Sprintf("Процесс %s создан %s(%s) в процессе %s (%s).",
		p.ProcessBefore.Name(),
		p.Who.Username, p.Who.Uid,
		p.WhoProcess.PID, p.WhoProcess.NameProcess)
}

func (p ProcessUpdate) Hash(hash hash.Hash) string {
	return hash(p.JSON())
}
