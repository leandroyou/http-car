package main

import (
	"context"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/go-redis/redis"
	_ "github.com/leandroyou/http-car/controller"
	"github.com/leandroyou/http-car/controller/route"
	_ "github.com/leandroyou/http-car/observer/listener"
	mongo2 "github.com/leandroyou/http-car/repository/mongo"
	redisRep "github.com/leandroyou/http-car/repository/redis"
	"github.com/leandroyou/http-car/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"os"
	"time"

	_ "github.com/leandroyou/http-car/docs"
	"github.com/swaggo/http-swagger"
)

// @title Simple HTTTP Server
// @version 1.0
// @description This is a simple http server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080/
// @BasePath
func main() {
	//Swagger
	route.Routes.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	//argparse
	parser := argparse.NewParser("Required-args", "Example of required vs optional arguments")

	// Optional shorthand argument
	mongoArg := parser.String("m", "mongo", &argparse.Options{Required: false, Help: "mongo is not a required option. serves to define the uri for the mongo database", Default: "mongodb"})
	redisArg := parser.String("r", "redis", &argparse.Options{Required: false, Help: "redis is not a required option. serves to define the uri for redis", Default: "redis"})

	//Ports
	mongoPortArg := parser.String("d", "mongoPort", &argparse.Options{Required: false, Help: "mongoPort is not a required option. serves to define the port for the mongo database", Default: "27017"})
	redisPortArg := parser.String("f", "redisPort", &argparse.Options{Required: false, Help: "redisPort is not a required option. serves to define the port for redis", Default: "6379"})

	// Parse args
	if err := parser.Parse(os.Args); err != nil {
		parser.Usage(err)
		panic(err)
	}

	//mongo - coldStorage
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//fmt.Sprintf("mongodb://%s:%s", *mongoArg, *mongoPortArg)
	//"mongodb://localhost:27017"
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", *mongoArg, *mongoPortArg)))
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
	//fmt.Sprintf("%s:%s", *redisArg, *redisPortArg)
	//"localhost:6379"
	rdbClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", *redisArg, *redisPortArg),
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

/*
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
*/