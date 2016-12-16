package router

import (
	"github.com/gorilla/mux"
	"net/http"
	task_handler "taskManagerWeb/handler"
)

func HandleRequest() {
	handler := mux.NewRouter()
	handler.HandleFunc("/tasks/save", task_handler.SaveTask).Methods(http.MethodPost)
	handler.HandleFunc("/tasks", task_handler.GetAllTask).Methods(http.MethodGet)
	handler.HandleFunc("/task/update", task_handler.UpdateTask).Methods(http.MethodPost)
	handler.HandleFunc("/task/delete", task_handler.DeleteTask).Methods(http.MethodPost)
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)
}
