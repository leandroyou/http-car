package car

import (
	"encoding/json"
	"github.com/leandroyou/http-car/middleware"
	"github.com/leandroyou/http-car/model/dao"
	"github.com/leandroyou/http-car/service"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// UpdateCar godoc
// @Summary Updates a car
// @Description Updates a car when a dao.Car is given
// @Accept  json
// @Produce  json
// @Param car body dao.Car true "Update car"
// @Success 200 {object} dao.Car
// @Failure 500 {object} dao.Car
// @Router /car/update [post]
func init() {
	http.Handle("/car/update", middleware.ProcessMiddleware(
		http.HandlerFunc(UpdateHandler),
	))
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		carToUpdate := dao.Car{}
		if err := json.Unmarshal(body, &carToUpdate); err != nil {
			panic(err)
		}

		if err := service.GetService().Storage.UpdateCar(r.Context(), carToUpdate); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{"status":"updated"}`)
		return
	}

	w.WriteHeader(http.StatusBadGateway)
	return
}
