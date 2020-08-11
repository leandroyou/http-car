package service

import (
	"context"
	"github.com/leandroyou/http-car/model/dao"
	"github.com/leandroyou/http-car/observer/event"
	"github.com/leandroyou/http-car/observer/subject"
	"github.com/leandroyou/http-car/repository"
)

type Storage struct {
}

func (Storage) CreateCar(ctx context.Context, car *dao.Car) error {
	if err := GetService().ColdStorage.CreateCar(ctx, car); err != nil {
		return err
	}
	go subject.Observers.Notify(event.Event{Car: *car})
	return nil
}

func (Storage) GetCars(ctx context.Context) ([]dao.Car, error) {
	return GetService().ColdStorage.GetCars(ctx)
}

func (Storage) GetCar(ctx context.Context, id string) (dao.Car, error) {
	hotCar, err := GetService().HotStorage.GetCar(ctx, id)
	if err == nil {
		return hotCar, nil
	}
	if err != repository.ErrNoCar {
		return dao.Car{}, err
	}

	coldCar, err := GetService().ColdStorage.GetCar(ctx, id)
	if err == repository.ErrNoCar {
		return dao.Car{}, nil
	}
	if err != nil {
		return dao.Car{}, err
	}
	//create goroutine
	go GetService().HotStorage.CreateCar(context.Background(), &coldCar)

	return coldCar, nil
}

func (Storage) DeleteCar(ctx context.Context, car dao.Car) error {
	if err := GetService().ColdStorage.DeleteCar(ctx, car); err != nil {
		return err
	}

	go GetService().HotStorage.DeleteCar(context.Background(), car)

	return nil
}

func (Storage) UpdateCar(ctx context.Context, car dao.Car) error {
	if err := GetService().ColdStorage.UpdateCar(ctx, car); err != nil {
		return err
	}

	go GetService().HotStorage.UpdateCar(context.Background(), car)
	return nil
}
