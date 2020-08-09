package main

import (
	"context"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	_ "github.com/leandroyou/http-car/controller"
	"github.com/leandroyou/http-car/controller/route"
	mongo2 "github.com/leandroyou/http-car/repository/mongo"
	redisRep "github.com/leandroyou/http-car/repository/redis"
	"github.com/leandroyou/http-car/service"
	"io"
	"net/http"
	"time"
)

func main() {
	//mongo - coldStorage
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		panic(err)
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	mongoRepository := mongo2.Repository{
		CarRepository: mongo2.CarRepository{Client: mongoClient},
	}

	//redis - hotStorage
	rdbClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := rdbClient.Ping().Result(); err != nil {
		panic(err)
	}

	redisRepository := redisRep.Repository{
		CarRepository: redisRep.CarRepository{Client: rdbClient},
	}

	srv := service.Service{
		ColdStorage: service.ColdStorage{
			IRepository: mongoRepository,
		},
		HotStorage: service.HotStorage{
			IRepository: redisRepository,
		},
		Storage: service.Storage{},
	}

	service.InitializeService(srv)

	// CreateCar an example endpoint/route
	/*http.Handle("/", middleware.ProcessMiddleware(
		http.HandlerFunc(ExampleHandler),
		middleware.GetRegisteredMiddlewares()...,
	))*/
	http.Handle("/", route.Routes)

	// Run...
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
