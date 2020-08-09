package car

import (
	"encoding/json"
	"fmt"
	"github.com/leandroyou/http-car/middleware"
	"github.com/leandroyou/http-car/model/dao"
	"github.com/leandroyou/http-car/service"
	"net/http"
)

func init() {
	http.Handle("/car/create", middleware.ProcessMiddleware(
		http.HandlerFunc(CreateHandler),
	))
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		newCar := dao.Car{}
		if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
			panic(err)
		}

		if err := service.GetService().Storage.CreateCar(r.Context(), &newCar); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"Id":"%s"}`, newCar.Id)
		return
	}

	w.WriteHeader(http.StatusBadGateway)
	return
}
