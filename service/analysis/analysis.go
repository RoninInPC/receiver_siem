package analysis

import (
	"receiver_siem/entity/subject/notification/receivernotification"
	"receiver_siem/storageuser"
	"receiver_siem/telegramsender"
	"slices"
	"strconv"
	"sync"
	"time"
)

type Analysis struct {
	channel chan receivernotification.Notification
	sync.Mutex
	notifications receivernotification.Notifications
	storageUsers  storageuser.StorageUser
	sender        telegramsender.TelegramSender[string]
	duration      time.Duration
}

func Init(
	channel chan receivernotification.Notification,
	storageUsers storageuser.StorageUser,
	sender telegramsender.TelegramSender[string],
	duration time.Duration) Analysis {
	return Analysis{channel: channel,
		Mutex:         sync.Mutex{},
		notifications: make(receivernotification.Notifications, 0),
		storageUsers:  storageUsers,
		sender:        sender,
		duration:      duration,
	}
}

func (a Analysis) Work() {
	go func() {
		for notification := range a.channel {
			a.Lock()
			a.notifications = append(a.notifications, notification)
			a.Unlock()
		}
	}()
	for {
		a.Lock()
		c := slices.Clone(a.notifications)
		a.notifications = make(receivernotification.Notifications, 0)
		a.Unlock()
		for _, message := range c.SortByHost().ToTelegramString() {
			for _, user := range a.storageUsers.GetUsers() {
				id, _ := strconv.ParseInt(user.UserId, 10, 64)
				a.sender.Send(id, message)
			}
		}

		time.Sleep(a.duration)
	}
}
