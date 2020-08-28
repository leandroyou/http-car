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

// DeleteCar godoc
// @Summary Deletes a car
// @Description Deletes a car when a dao.Car.Id is given
// @Accept  json
// @Produce  json
// @Param car body dao.Car.Id true "Delete a car"
// @Success 200 {string} Car deleted
// @Failure 500 {string} Error
// @Router /car/delete [post]
func init() {
	http.Handle("/car/delete", middleware.ProcessMiddleware(
		http.HandlerFunc(DeleteHandler),
	))
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		carToDelete := dao.Car{}
		if err := json.Unmarshal(body, &carToDelete); err != nil {
			panic(err)
		}

		if err := service.GetService().Storage.DeleteCar(r.Context(), carToDelete); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{"status":"deleted"}`)
		return
	}

	w.WriteHeader(http.StatusBadGateway)
	return
}
