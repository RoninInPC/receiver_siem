package receivernotification

import (
	"encoding/json"
	"fmt"
	"receiver_siem/entity/subject"
	"receiver_siem/entity/subject/notification"
	"receiver_siem/hostinfo"
	"strings"
	"time"
)

const (
	FileChangeT    subject.SubjectType = 101
	FileDeleteT    subject.SubjectType = 102
	FileNewT       subject.SubjectType = 103
	FileRenameT    subject.SubjectType = 104
	PortDeleteT    subject.SubjectType = 105
	PortNewT       subject.SubjectType = 106
	PortUpdateT    subject.SubjectType = 107
	ProcessDeleteT subject.SubjectType = 108
	ProcessNewT    subject.SubjectType = 109
	ProcessUpdateT subject.SubjectType = 110
	ProcessEndT    subject.SubjectType = 111
	UserDeleteT    subject.SubjectType = 112
	UserNewT       subject.SubjectType = 113
	UserUpdateT    subject.SubjectType = 114
)

type GetBaseNotificationInfo interface {
	GetHostName() string
	GetUsername() string
	GetProcessPID() string
	GetTime() time.Time
	SetUser(user *subject.User)
	SetProcess(process *subject.Process)
	GetProcessInfo() string
	GetProcessInfoMarkdown() string
	GetUserInfo() string
	GetUserInfoMarkdown() string
}

type Notification interface {
	notification.Notification
	GetBaseNotificationInfo
	GetInfo() string
	GetInfoMarkdown() string
}

type BaseNotification struct {
	Who        *subject.User
	WhoProcess *subject.Process
	Time       time.Time
	Host       hostinfo.HostInfo
}

func (b BaseNotification) GetProcessInfo() string {
	return fmt.Sprintf(
		"Name: %s;\n"+
			"PID: %s;\n"+
			"Nice: %d;\n"+
			"IsRunning: %t;\n"+
			"IsBackground: %t;\n"+
			"CreateTime: %s;\n"+
			"CMDLine: %s;\n",
		b.WhoProcess.NameProcess,
		b.WhoProcess.PID,
		b.WhoProcess.Nice,
		b.WhoProcess.IsRunning,
		b.WhoProcess.IsBackGround,
		b.WhoProcess.CreateTime.Format("2006-01-02 15:04:05"),
		b.WhoProcess.CMDLine)
}

func (b BaseNotification) GetProcessInfoMarkdown() string {
	return fmt.Sprintf(
		"*Name:* %s;\n"+
			"*PID:* %s;\n"+
			"*Nice:* %d;\n"+
			"*IsRunning:* %t;\n"+
			"*IsBackground:* %t;\n"+
			"*CreateTime:* %s;\n"+
			"*CMDLine:* %s;\n",
		b.WhoProcess.NameProcess,
		b.WhoProcess.PID,
		b.WhoProcess.Nice,
		b.WhoProcess.IsRunning,
		b.WhoProcess.IsBackGround,
		b.WhoProcess.CreateTime.Format("2006-01-02 15:04:05"),
		b.WhoProcess.CMDLine)
}

func (b BaseNotification) GetUserInfo() string {
	return fmt.Sprintf(
		"UserName: %s;\n"+
			"UID: %s;\n"+
			"GID: %s;\n"+
			"HomeDir: %s;\n",
		b.Who.Username,
		b.Who.Uid,
		b.Who.Gid,
		b.Who.HomeDir)
}

func (b BaseNotification) GetUserInfoMarkdown() string {
	return fmt.Sprintf(
		"*UserName:* %s;\n"+
			"*UID:* %s;\n"+
			"*GID:* %s;\n"+
			"*HomeDir:* %s;\n",
		b.Who.Username,
		b.Who.Uid,
		b.Who.Gid,
		b.Who.HomeDir)
}

func (b BaseNotification) GetHostName() string {
	return b.Host.HostName
}

func (b BaseNotification) GetUsername() string {
	return b.Who.Username
}

func (b BaseNotification) GetProcessPID() string {
	return b.WhoProcess.PID
}

func (b BaseNotification) GetTime() time.Time {
	return b.Time
}

func (b BaseNotification) SetUser(user *subject.User) {
	b.Who = user
}

func (b BaseNotification) SetProcess(process *subject.Process) {
	b.WhoProcess = process
}

func unmarshal[A any](jsoned string) A {
	var a A
	json.Unmarshal([]byte(jsoned), &a)
	return a
}

func JsonedToNotification(jsoned string, subjectType subject.SubjectType) Notification {
	jsoned = strings.Replace(jsoned, "\\\"", "\"", -1)
	switch subjectType {
	case FileChangeT:
		return unmarshal[FileUpdate](jsoned)
	case FileDeleteT:
		return unmarshal[FileDelete](jsoned)
	case FileNewT:
		return unmarshal[FileNew](jsoned)
	case FileRenameT:
		return unmarshal[FileRename](jsoned)
	case PortDeleteT:
		return unmarshal[PortDelete](jsoned)
	case PortNewT:
		return unmarshal[PortNew](jsoned)
	case PortUpdateT:
		return unmarshal[PortUpdate](jsoned)
	case ProcessDeleteT:
		return unmarshal[ProcessDelete](jsoned)
	case ProcessNewT:
		return unmarshal[ProcessNew](jsoned)
	case ProcessUpdateT:
		return unmarshal[ProcessUpdate](jsoned)
	case ProcessEndT:
		return unmarshal[ProcessEnd](jsoned)
	case UserDeleteT:
		return unmarshal[UserDelete](jsoned)
	case UserNewT:
		return unmarshal[UserNew](jsoned)
	case UserUpdateT:
		return unmarshal[UserUpdate](jsoned)
	}
	return nil
}

type Notifications []Notification

func (n Notifications) Len() int {
	return len(n)
}

func (n Notifications) Less(i, j int) bool {
	return n[i].GetTime().Before(n[j].GetTime()) && n[i].GetHostName() < n[j].GetHostName()
}

func (n Notifications) Swap(i, j int) {
	c := n[i]
	n[i] = n[j]
	n[j] = c
}

type NotificationsByHost map[string]Notifications

func (n Notifications) SortByHost() NotificationsByHost {
	answer := make(map[string]Notifications)
	for _, part := range n {
		if _, ok := answer[part.GetHostName()]; !ok {
			answer[part.GetHostName()] = make(Notifications, 0)
		} else {
			answer[part.GetHostName()] = append(answer[part.GetHostName()], part)
		}
	}
	return answer
}

func (n NotificationsByHost) ToString() string {
	answer := ""
	for host, notifications := range n {
		answer += host + ":\n"
		for _, n := range notifications {
			answer += "\t" + n.Name() + ";\n"
		}
	}
	return answer
}

const (
	TelegramMessageConst = 4096
)

func (n NotificationsByHost) ToTelegramString() []string {
	answer := make([]string, 0)
	part := ""
	for host, notifications := range n {
		part += host + ":\n"
		for _, n := range notifications {
			sum := "\t" + n.Name() + ";\n"
			if len(part)+len(sum)+6 > TelegramMessageConst {
				answer = append(answer, part)
				part = sum
			} else {
				part += sum
			}
		}
		answer = append(answer, part)
	}
	for i, p := range answer {
		switch i {
		case 0:
			p += "..."
			break
		case len(answer) - 1:
			p = "..." + p
			break
		default:
			p = "..." + p + "..."
			break
		}
	}
	return answer
}
