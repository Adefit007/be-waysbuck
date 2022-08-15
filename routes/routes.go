package routes

//import gorilla mux
import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	ProductRoutes(r)
	UserRoutes(r)
	ProfileRoutes(r)
	AuthRoutes(r)
	ToppingRoutes(r)
}
