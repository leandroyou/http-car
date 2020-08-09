package route

import "github.com/gorilla/mux"

var Routes = mux.NewRouter()

func init() {
	Routes.StrictSlash(true)
}
