package handler

import (
	"net/http"
	"strings"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"taskManagerClient/contract"
	"taskManagerWeb/model"
	"bytes"
)

func SaveTask(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	task := strings.Join(req.Form["task"], "")
	data := &contract.Task{}
	data.Task = &task
	dataToSend, err := proto.Marshal(data)
	if (err != nil) {
		log.Fatal("error occurs while creationg contract.")
		return
	}

	request, err := model.CreateRequest(http.MethodPost, "http://localhost:3000/tasks", bytes.NewBuffer(dataToSend))
	if (err != nil) {
		log.Fatalln("got error while creating server..")
		return
	}
	client := http.Client{};
	_, err = client.Do(request);
	if (err != nil) {
		log.Fatalln("got error while calling server;..")
		return
	}
	res.Write([]byte("task has stored"))
}

func GetAllTask(res http.ResponseWriter, req *http.Request) {
	fmt.Println("edcw")
	request, _ := model.CreateRequest(http.MethodGet, "http://localhost:3000/tasks", nil)
	client := http.Client{}
	response, err := client.Do(request)
	if (err != nil) {
		log.Fatalln("got error while calling server;..")
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	res.Write([]byte(body))
}