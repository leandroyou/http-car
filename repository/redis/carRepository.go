package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/leandroyou/http-car/model/dao"
	"github.com/leandroyou/http-car/repository"
)

type CarRepository struct {
	*redis.Client
}

func (r CarRepository) CreateCar(ctx context.Context, car *dao.Car) error {
	jsonCar, err := json.Marshal(car)
	if err != nil {
		return err
	}

	err = r.Client.Set(car.Id, jsonCar, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r CarRepository) GetCars(ctx context.Context) ([]dao.Car, error) {
	return nil, nil
}

func (r CarRepository) GetCar(ctx context.Context, id string) (dao.Car, error) {
	val, err := r.Client.Get(id).Result()
	if err != nil {
		if err == redis.Nil {
			return dao.Car{}, repository.ErrNoCar
		}
		return dao.Car{}, err
	}

	var car dao.Car
	if err := json.Unmarshal([]byte(val), &car); err != nil {
		return dao.Car{}, err
	}

	return car, nil
}

func (r CarRepository) DeleteCar(ctx context.Context, car dao.Car) error {
	return r.Client.Del(car.Id).Err()
}

func (r CarRepository) UpdateCar(ctx context.Context, car dao.Car) error {
	return r.CreateCar(ctx, &car)
}
