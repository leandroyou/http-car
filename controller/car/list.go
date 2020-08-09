package car

import (
	"encoding/json"
	"fmt"
	"github.com/leandroyou/http-car/controller/route"
	"github.com/leandroyou/http-car/service"
	"net/http"
)

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
