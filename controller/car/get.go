package car

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leandroyou/http-car/controller/route"
	_ "github.com/leandroyou/http-car/model/dao"
	"github.com/leandroyou/http-car/service"
	"net/http"
)

// GetCar godoc
// @Summary Get a car
// @Description Get a car when an Id is given
// @Accept  json
// @Produce  json
// @Param id path string true "Car ID"
// @Success 200 {object} dao.Car
// @Failure 500 {string} Error
// @Router /car/get/{id} [get]
func init() {
	route.Routes.HandleFunc("/car/get/{id}", GetHandler)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	car, err := service.GetService().Storage.GetCar(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	//TODO: Converter para dto

	carsBytes, err := json.Marshal(car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprint(w, string(carsBytes)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
