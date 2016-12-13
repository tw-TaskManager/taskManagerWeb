package handler

import (
	"net/http"
	"strings"
	"fmt"
	"log"
	"io/ioutil"
)

func SaveTask(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	task := strings.Join(req.Form["task"], "")
	fmt.Println(task)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:3000/tasks", strings.NewReader(task))
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
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/tasks", nil)

	client := http.Client{}
	response, err := client.Do(request)
	if (err != nil) {
		log.Fatalln("got error while calling server;..")
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	res.Write([]byte(body))
}