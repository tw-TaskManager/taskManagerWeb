package router

import (
	"github.com/gorilla/mux"
	"net/http"
	task_handler "taskManagerWeb/handler"
)

func HandleRequest() {
	handler := mux.NewRouter()
	handler.HandleFunc("/tasks", task_handler.SaveTask).Methods(http.MethodPost)
	handler.HandleFunc("/tasks", task_handler.GetAllTask).Methods(http.MethodGet)
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)
}
