package main

import (
	"day15/Item/routers"
	"log"
	"net/http"
)

func main() {

	router := routers.InitRouters()

	log.Fatal(http.ListenAndServe(":8888", router))

}
