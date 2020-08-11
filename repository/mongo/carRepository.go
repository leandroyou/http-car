package mongo

import (
	"context"
	"fmt"
	"github.com/leandroyou/http-car/model/dao"
	"github.com/leandroyou/http-car/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CarRepository struct {
	*mongo.Client
}

func (r CarRepository) CreateCar(ctx context.Context, car *dao.Car) error {
	collection := r.Client.Database("testCar").Collection("car")

	res, err := collection.InsertOne(ctx, car)
	if err != nil {
		return err
	}

	car.Id = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (r CarRepository) GetCars(ctx context.Context) ([]dao.Car, error) {
	collection := r.Client.Database("testCar").Collection("car")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var cars []dao.Car
	if err := cursor.All(ctx, &cars); err != nil {
		return nil, err
	}

	return cars, nil
}

func (r CarRepository) GetCar(ctx context.Context, id string) (dao.Car, error) {
	collection := r.Client.Database("testCar").Collection("car")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dao.Car{}, err
	}
	res := collection.FindOne(ctx, bson.M{"_id": oid})

	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return dao.Car{}, repository.ErrNoCar
		}
		return dao.Car{}, err
	}

	var car dao.Car
	if err := res.Decode(&car); err != nil {
		return dao.Car{}, err
	}

	return car, nil
}

func (r CarRepository) DeleteCar(ctx context.Context, car dao.Car) error {
	collection := r.Client.Database("testCar").Collection("car")

	oid, err := primitive.ObjectIDFromHex(car.Id)
	if err != nil {
		return err
	}
	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	fmt.Println("DeleteOne Result:", res)

	return nil
}

func (r CarRepository) UpdateCar(ctx context.Context, car dao.Car) error {
	collection := r.Client.Database("testCar").Collection("car")

	oid, err := primitive.ObjectIDFromHex(car.Id)
	if err != nil {
		return err
	}
	res, err := collection.UpdateOne(ctx,
		bson.M{"_id": oid},
		bson.M{
			"$set": car,
		})
	if err != nil {
		return err
	}

	fmt.Println("UpdatedOne Result:", res)

	return nil
}
