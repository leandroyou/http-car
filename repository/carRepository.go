package repository

import (
	"context"
	"github.com/leandroyou/http-car/model/dao"
)

type ICarRepository interface {
	CreateCar(ctx context.Context, car *dao.Car) error
	GetCars(ctx context.Context) ([]dao.Car, error)
	GetCar(ctx context.Context, id string) (dao.Car, error)
	DeleteCar(ctx context.Context, car dao.Car) error
	UpdateCar(ctx context.Context, car dao.Car) error
}
