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
		_, err = w.Write([]byte("\n\n\n"))
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
	var body interface{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		messages = append(messages, err.Error())
	}

	switch body.(type) {
	case map[string]interface{}:
		messages = append(messages, "map[string]interface{}")
	case interface{}:
		messages = append(messages, "map[string]interface{}")
	case map[string]string:
		messages = append(messages, "map[string]string")
	case map[string]json.RawMessage:
		messages = append(messages, "map[string]json.RawMessage")
	default:
		messages = append(messages, "unknown")
	}

	messages = append(messages, fmt.Sprint(body))
	//log.Println(m)
	api := ServiceData{}

	alertData := model.T{}
	if data, ok := body.(map[string]interface{})["alert_data"]; ok {
		messages = append(messages, fmt.Sprint(data))
		b, err := json.Marshal(data)
		if err != nil {
			messages = append(messages, err.Error())
		}
		err = json.Unmarshal(b, &alertData)
		if err != nil {
			messages = append(messages, err.Error())
		}
		api.APIUrl = alertData.Url.Full
		api.StatusCode = alertData.Http.Response.StatusCode
		api.ServiceName = alertData.Service.Name
	} else {
		messages = append(messages, "data, ok := body.(map[string]interface{})[\"alert_data\"]; ok ")
	}

	str := fmt.Sprint(api.APIUrl, " , ", api.StatusCode, " , ", api.ServiceName)

	//if len(messages) > 5 {
	//	messages = make([]string, 0)
	//}
	messages = append(messages, str)
	_, err = w.Write([]byte(str))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	} else {
		log.Println("hello")
	}
}
