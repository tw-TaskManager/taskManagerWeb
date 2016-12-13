package main

import (
	"taskManagerWeb/router"
	"net/http"
)

func main() {

	router.HandleRequest()
	http.ListenAndServe(":4000", nil)
}