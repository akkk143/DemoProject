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

// Start starts the http server
func Start() {
	messages = append(messages, "started...")
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
	ServiceName    string
	TransactionUrl string
	StatusCode     int
	APIUrl         string
}

func customQueryAlert(w http.ResponseWriter, r *http.Request) {
	var body interface{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		messages = append(messages, err.Error())
	}

	messages = append(messages, fmt.Sprint(body))

	data := body.(map[string]interface{})["alert_data"]

	api := ServiceData{}

	alertData := model.AlertData{}
	if d, ok := data.(string); ok {

		err = json.Unmarshal([]byte(d), &alertData)
		if err != nil {
			messages = append(messages, err.Error())
		}
		api.TransactionUrl = alertData.Transaction.Name
	} else {
		messages = append(messages, "could not convert data to string")
	}
	str := fmt.Sprint(api.APIUrl, " , ", api.StatusCode, " , ", api.ServiceName)
	messages = append(messages, str)
	_, err = w.Write([]byte(str))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	} else {
		log.Println("hello")
	}
}
