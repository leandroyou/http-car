package car

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leandroyou/http-car/controller/route"
	"github.com/leandroyou/http-car/service"
	"net/http"
)

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
