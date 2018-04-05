package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.InitRouters {
	router := mux.NewRouter().StrictSlash(false)

	router = setSellingRouter(router)

	return router
}
