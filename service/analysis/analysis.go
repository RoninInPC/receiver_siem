package analysis

import (
	"receiver_siem/entity/subject/notification/receivernotification"
	"sort"
	"sync"
	"time"
)

type Notifications []receivernotification.Notification

func (n Notifications) Len() int {
	return len(n)
}

func (n Notifications) Less(i, j int) bool {
	return n[i].GetTime().Before(n[j].GetTime())
}

func (n Notifications) Swap(i, j int) {
	c := n[i]
	n[i] = n[j]
	n[j] = c
}

type Analysis struct {
	channel chan receivernotification.Notification
	sync.Mutex
	notifications Notifications
	duration      time.Duration
}

func Init(channel chan receivernotification.Notification, duration time.Duration) *Analysis {
	return &Analysis{channel: channel,
		Mutex:         sync.Mutex{},
		notifications: make(Notifications, 0),
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
		c := make(Notifications, len(a.notifications))
		copy(a.notifications, c)
		a.notifications = make(Notifications, 0)
		a.Unlock()

		sort.Sort(c)

		time.Sleep(a.duration)
	}
}
