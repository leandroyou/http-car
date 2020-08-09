package service

import (
	"context"
	"github.com/leandroyou/http-car/model/dao"
	"github.com/leandroyou/http-car/repository"
)

type HotStorage struct {
	repository.IRepository
}

func (s HotStorage) createCar(ctx context.Context, car *dao.Car) error {
	return s.IRepository.CreateCar(ctx, car)
}

func (s HotStorage) getCars(ctx context.Context) ([]dao.Car, error) {
	return s.IRepository.GetCars(ctx)
}

func (s HotStorage) GetCar(ctx context.Context, id string) (dao.Car, error) {
	return s.IRepository.GetCar(ctx, id)
}

func (s HotStorage) DeleteCar(ctx context.Context, car dao.Car) error {
	return s.DeleteCar(ctx, car)
}

func (s HotStorage) UpdateCar(ctx context.Context, car dao.Car) error {
	return s.IRepository.UpdateCar(ctx, car)
}
