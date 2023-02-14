package app

import (
	"DemoProject/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const serviceRoute = "/demo/project"

var srv *http.Server
var messages []string
var count int

// Start starts the http server
func Start() {
	messages = append(messages, "Hello World!")
	count = 0
	createServer()
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		panic(err)
	}
}

func createServer() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc(serviceRoute+"/printmsg", printMsg).Methods("GET")
	r.HandleFunc(serviceRoute+"/customquery/alert", customQueryAlert).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func printMsg(w http.ResponseWriter, r *http.Request) {
	for _, m := range messages {
		_, err := w.Write([]byte(m))
		if err != nil {
			log.Printf("couldnt write response error [%s]\n", err)
		}
	}
}

type ServiceData struct {
	ServiceName string
	APIUrl      string
	StatusCode  int
}

func customQueryAlert(w http.ResponseWriter, r *http.Request) {
	body := model.AlertData{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		messages = append(messages, err.Error())
	}

	if err != nil {
		messages = append(messages, err.Error())
	}
	messages = append(messages, fmt.Sprint(body))
	api := ServiceData{}

	api.ServiceName = body.Service.Name
	api.StatusCode = body.Http.Response.StatusCode
	api.APIUrl = body.Url.Full

	str := fmt.Sprint(api.APIUrl, " , ", api.StatusCode, " , ", api.ServiceName)

	if len(messages) > 5 {
		messages = make([]string, 0)
	}
	messages = append(messages, str)
	_, err = w.Write([]byte(str))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	} else {
		log.Println("hello")
	}
}
