package main

import (
	"task_manager_web/router"
	"net/http"
)

func CreateTcpConnection() {

}

func main() {

	router.HandleRequest()
	http.ListenAndServe(":4000", nil)
}