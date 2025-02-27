package receivernotification

import (
	"encoding/json"
	"fmt"
	"receiver_siem/entity/subject"
	"receiver_siem/hash"
)

type PortUpdate struct {
	PortBefore subject.PortTables
	PortAfter  subject.PortTables
	BaseNotification
}

func (p PortUpdate) GetInfo() string {
	//TODO implement me
	panic("implement me")
}

func (p PortUpdate) GetInfoMarkdown() string {
	//TODO implement me
	panic("implement me")
}

func (p PortUpdate) JSON() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (p PortUpdate) Type() subject.SubjectType {
	return PortUpdateT
}

func (p PortUpdate) Name() string {
	return fmt.Sprintf("Порт %s изменён %s(%s) в процессе %s (%s).",
		p.PortBefore.Name(),
		p.Who.Username, p.Who.Uid,
		p.WhoProcess.PID, p.WhoProcess.NameProcess)
}

func (p PortUpdate) Hash(hash hash.Hash) string {
	return hash(p.JSON())
}
