package main

import (
	"armeris/gorest/routers"
	"log"
	"net/http"
)

func main() {

	router := routers.NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
