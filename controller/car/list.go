package car

import (
	"encoding/json"
	"fmt"
	"github.com/leandroyou/http-car/controller/route"
	_ "github.com/leandroyou/http-car/model/dao"
	"github.com/leandroyou/http-car/service"
	"net/http"
)

// GetCars godoc
// @Summary Get all cars
// @Description Get all the cars
// @Accept  json
// @Produce  json
// @Success 200 {array} dao.Car
// @Failure 500 {string} Error
// @Router /car/get [get]
func init() {
	route.Routes.HandleFunc("/car/get", ListHandler)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	cars, err := service.GetService().Storage.GetCars(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	//TODO: Converter para dto

	carsBytes, err := json.Marshal(cars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprint(w, string(carsBytes)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
