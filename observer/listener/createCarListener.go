package listener

import (
	"fmt"
	"github.com/leandroyou/http-car/observer/event"
	"github.com/leandroyou/http-car/observer/subject"
)

type EventObserverGoRoutine struct {
	Description string
}

func (e *EventObserverGoRoutine) NotifyCallback(event event.Event) {
	fmt.Println("eventObserverGoRoutine")
	//service := (interface{}(&event.Service)).(teste.Service)
	//service.HotStorage.CreateCar(context.Background(), &event.Car)
}

func init(){
	var goroutineObserver = EventObserverGoRoutine{Description: "goroutine observer"}
	subject.Observers.AddListener(&goroutineObserver)
}
