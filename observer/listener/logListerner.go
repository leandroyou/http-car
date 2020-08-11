package listener

import (
	"fmt"
	"github.com/leandroyou/http-car/observer/event"
	"github.com/leandroyou/http-car/observer/subject"
)

type EventObserverLogs struct {
	Description string
}

func (e *EventObserverLogs) NotifyCallback(event event.Event) {
	fmt.Println("Car created: ", event.Car.Id)
}

func init(){
	var logObserver = EventObserverLogs{Description: "Log observer"}
	subject.Observers.AddListener(&logObserver)
}