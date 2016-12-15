package handler

import (
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"taskManagerClient/contract"
	"taskManagerWeb/model"
	"bytes"
	"fmt"
	"strconv"
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

	request, err := model.CreateRequest(http.MethodPost, "http://localhost:3000/tasks/save", bytes.NewBuffer(dataToSend))
	if (err != nil) {
		log.Fatalln("got error while creating server..")
		return
	}
	client := http.Client{};
	response, err := client.Do(request);
	body, err := ioutil.ReadAll(response.Body)
	contractOfResponse := contract.Response{}
	err = proto.Unmarshal(body, &contractOfResponse)
	if (err != nil) {
		log.Fatalln("got error while calling server;..")
		return
	}
	res.Write(contractOfResponse.Response)
}

func UpdateTask(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	content := strings.Join(req.Form["content"], "")
	taskId := req.Form["id"][0];
	id64, err := strconv.ParseInt(taskId, 10, 32)
	id32 := int32(id64)
	data := &contract.Task{}
	data.Task = &content
	data.Id = &id32
	dataToSend, err := proto.Marshal(data)
	if (err != nil) {
		log.Fatal("error occurs while creationg contract.")
		return
	}

	request, err := model.CreateRequest(http.MethodPost, "http://localhost:3000/task/update", bytes.NewBuffer(dataToSend))
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
	res.Write([]byte("task has updated"))
}

func GetAllTask(res http.ResponseWriter, req *http.Request) {
	request, _ := model.CreateRequest(http.MethodGet, "http://localhost:3000/tasks", nil)
	client := http.Client{}
	response, err := client.Do(request)
	if (err != nil) {
		log.Fatalln("got error while calling server;..")
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	contractOfResponse := contract.Response{}
	err = proto.Unmarshal(body, &contractOfResponse)
	fmt.Println()
	if (err != nil) {
		log.Fatalln("got error while parsing;..")
		return
	}
	res.Write(contractOfResponse.Response)
}