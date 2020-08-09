package service

import (
	"context"
	"github.com/leandroyou/http-car/model/dao"
	"github.com/leandroyou/http-car/repository"
)

type ColdStorage struct {
	repository.IRepository
}

func (s ColdStorage) CreateCar(ctx context.Context, car *dao.Car) error {
	return s.IRepository.CreateCar(ctx, car)
}

func (s ColdStorage) GetCars(ctx context.Context) ([]dao.Car, error) {
	return s.IRepository.GetCars(ctx)
}

func (s ColdStorage) GetCar(ctx context.Context, id string) (dao.Car, error) {
	return s.IRepository.GetCar(ctx, id)
}

func (s ColdStorage) DeleteCar(ctx context.Context, car dao.Car) error {
	return s.IRepository.DeleteCar(ctx, car)
}

func (s ColdStorage) UpdateCar(ctx context.Context, car dao.Car) error {
	return s.IRepository.UpdateCar(ctx, car)
}
