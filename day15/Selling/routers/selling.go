package routers

import (
	"day15//Selling/controllers"

	"github.com/gorilla/mux"
)

func setSellingRouters(router *mux.Router) *mux.Router {

	router.HandleFunc("/selling", controllers.GetSelling).Methods("GET")
	return router
}
