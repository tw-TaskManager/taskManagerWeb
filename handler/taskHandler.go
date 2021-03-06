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
	"strconv"
	"time"
)

func SaveTask(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("taskManager")
	req.ParseForm()
	task := strings.Join(req.Form["task"], "")
	data := &contract.Task{}
	data.Task = &task
	dataToSend, err := proto.Marshal(data)
	if (err != nil) {
		log.Println(err.Error())
		return
	}
	taskRequest := "http://localhost:3000/tasks/save/" + cookie.Value
	request, err := model.CreateRequest(http.MethodPost, taskRequest, bytes.NewBuffer(dataToSend))
	if (err != nil) {
		log.Println(err.Error())
		return
	}
	client := http.Client{};
	response, err := client.Do(request);
	body, err := ioutil.ReadAll(response.Body)
	if (err != nil) {
		log.Println("got error while reading")
		return
	}
	contractOfResponse := contract.Response{}
	err = proto.Unmarshal(body, &contractOfResponse)
	if (err != nil) {
		log.Println("got error while parsing task")
		return
	}
	res.Write(contractOfResponse.Response)
}

func UpdateTask(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("taskManager")
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
		log.Println(err.Error())
		return
	}
	taskRequest := "http://localhost:3000/task/update/" + cookie.Value
	request, err := model.CreateRequest(http.MethodPost, taskRequest, bytes.NewBuffer(dataToSend))
	if (err != nil) {
		log.Println(err.Error())
		return
	}
	client := http.Client{};
	_, err = client.Do(request);
	if (err != nil) {
		log.Println("got error while calling server; http://localhost:3000/task/update/" + cookie.Value)
		return
	}
	res.Write([]byte("task has updated"))
}

func GetAllTask(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("taskManager")
	if (err != nil) {
		http.Redirect(res, req, "/login.html", http.StatusTemporaryRedirect)
		return
	}
	taskRequest := "http://localhost:3000/tasks/" + cookie.Value
	request, _ := model.CreateRequest(http.MethodGet, taskRequest, nil)
	client := http.Client{}
	response, err := client.Do(request)
	if (err != nil) {
		log.Println("got error while calling server; http://localhost:3000/tasks/" + cookie.Value)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	contractOfResponse := contract.Response{}
	err = proto.Unmarshal(body, &contractOfResponse)
	if (err != nil) {
		return
	}
	res.Write(contractOfResponse.Response)
}

func DeleteTask(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	cookie, err := req.Cookie("taskManager")
	taskId := req.Form["id"][0];
	id64, err := strconv.ParseInt(taskId, 10, 32)
	id32 := int32(id64)
	data := &contract.Task{}
	data.Id = &id32
	dataToSend, err := proto.Marshal(data)
	if (err != nil) {
		log.Println(err.Error())
		return
	}
	taskRequest := "http://localhost:3000/task/delete/" + cookie.Value
	request, err := model.CreateRequest(http.MethodPost, taskRequest, bytes.NewBuffer(dataToSend))
	if (err != nil) {
		log.Println("got error while creating server; http://localhost:3000/tasks/delete/" + cookie.Value)
		return
	}
	client := http.Client{};
	_, err = client.Do(request);
	if (err != nil) {
		log.Println("got error while calling server;..")
		return
	}
	res.Write([]byte("deleted."))
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	userName := req.Form["userName"][0]
	emailId := req.Form["emailId"][0]
	password := req.Form["password"][0]
	user := &contract.User{}
	user.UserName = &userName
	user.EmailId = &emailId
	user.Password = &password
	data_to_send, err := proto.Marshal(user)
	if (err != nil) {
		log.Println("error occurs while creationg contract for user.")
		return
	}
	request, err := model.CreateRequest(http.MethodPost, "http://localhost:5000/task/createUser", bytes.NewBuffer(data_to_send))
	if (err != nil) {
		log.Println("got error while creating server; http://localhost:3000/tasks/createUser")
		return
	}
	client := http.Client{};
	response, err := client.Do(request);
	if (err != nil) {
		log.Println("got error while calling server.....")
		return
	}

	if (response.StatusCode == http.StatusConflict) {
		res.Write([]byte("user is already exist"))
		return
	}

	http.Redirect(res, req, "/login.html", http.StatusTemporaryRedirect)

}

func Auth(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	emailId := req.Form["emailId"][0]
	password := req.Form["password"][0]
	user := &contract.User{}
	user.EmailId = &emailId
	user.Password = &password
	data_to_send, err := proto.Marshal(user)
	if (err != nil) {
		log.Println("error occurs while creationg contract for user.")
		return
	}

	request, err := model.CreateRequest(http.MethodPost, "http://localhost:5000/task/login", bytes.NewBuffer(data_to_send))
	client := http.Client{}
	response, err := client.Do(request)
	if (response.StatusCode == http.StatusForbidden) {
		res.Write([]byte("user not found"))
		return
	}
	cookies := http.Cookie{
		Name:"taskManager",
		Value:response.Cookies()[0].Value,
		Path:"/",
	}
	http.SetCookie(res, &cookies)
	http.Redirect(res, req, "/", http.StatusMovedPermanently)
	return
}

func Logout(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	cookies := http.Cookie{
		Name:"taskManager",
		Value:"",
		Path:"/",
		Expires:time.Now().AddDate(-10, -4, -1),
	}
	request, _ := model.CreateRequest(http.MethodGet, "http://localhost:5000/task/logout", nil)
	client := http.Client{}
	http.SetCookie(res, &cookies)
	client.Do(request)
	res.Write([]byte("/login.html"))
	return
}

func UserAlreadyLogin(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("taskManager")
	if (err != nil) {
		log.Println(err.Error())
		http.Redirect(res, req, "/login.html", http.StatusMovedPermanently)
	}
	if (cookie.Value != "") {
		log.Println("refdirecting to task page " + cookie.Value)
		http.Redirect(res, req, "/", http.StatusMovedPermanently)
		return
	}

}
