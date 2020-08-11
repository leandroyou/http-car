package event

import (
	"github.com/leandroyou/http-car/model/dao"
)

type Event struct {
	//Storage service.Storage
	Car dao.Car
}
