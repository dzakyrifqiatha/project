package main

import (
	"day15/Officer/routers"
	"log"
	"net/http"
)

func main() {

	router := routers.InitRouters()

	log.Fatal(http.ListenAndServe(":8887", router))

}
