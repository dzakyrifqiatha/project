package routers

import (
	"day15/Item/controllers"

	"github.com/gorilla/mux"
)

func setItemRouters(router *mux.Router) *mux.Router {

	router.HandleFunc("/item", controllers.GetItem).Methods("GET")
	return router
}
