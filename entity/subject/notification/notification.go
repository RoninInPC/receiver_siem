package notification

import "receiver_siem/entity/subject"

type Notification interface {
	subject.Subject
}
