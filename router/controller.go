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
	handler.HandleFunc("/task/createUser", task_handler.CreateUser).Methods(http.MethodPost)
	handler.HandleFunc("/task/login", task_handler.Auth).Methods(http.MethodPost)
	handler.HandleFunc("/task/logout", task_handler.Logout).Methods(http.MethodGet)
	handler.HandleFunc("/task/isUserAlreadyLogin",task_handler.UserAlreadyLogin).Methods(http.MethodGet)
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)
}
