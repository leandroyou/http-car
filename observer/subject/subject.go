package subject

import (
	"github.com/leandroyou/http-car/observer/event"
	"sync"
)

type Observer interface {
	NotifyCallback(event event.Event)
}

type Subject interface {
	AddListener(Observer)
	RemoveListener(Observer)
	Notify(event event.Event)
}

type EventSubject struct {
	Observers sync.Map
}

func (s *EventSubject) AddListener(obs Observer) {
	s.Observers.Store(obs, struct{}{})
}

func (s *EventSubject) RemoveListener(obs Observer) {
	s.Observers.Delete(obs)
}

func (s *EventSubject) Notify(event event.Event) {
	s.Observers.Range(func(key interface{}, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}

		key.(Observer).NotifyCallback(event)
		return true
	})
}

var Observers = EventSubject{
	Observers: sync.Map{},
}
