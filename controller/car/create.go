package car

import (
	"encoding/json"
	"fmt"
	"github.com/leandroyou/http-car/middleware"
	"github.com/leandroyou/http-car/model/dao"
	"github.com/leandroyou/http-car/service"
	"net/http"
)

// CreateCar godoc
// @Summary Create a car
// @Description Creates a new car when a dao.Car is given
// @Accept  json
// @Produce  json
// @Param car body dao.Car true "Add car"
// @Success 200 {object} dao.Car
// @Failure 500 {object} dao.Car
// @Router /car/create [post]
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
		if _, err := fmt.Fprintf(w, `{"Id":"%s"}`, newCar.Id); err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusBadGateway)
	return
}
