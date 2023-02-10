package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const serviceRoute = "/demo/project"

var srv *http.Server
var messages string

type message struct {
	ServiceName     string `json:"service_name"`
	Reason          string `json:"reason"`
	TransactionType string `json:"transaction_type"`
}

// Start starts the http server
func Start() {
	messages = "Hello World!"
	createServer()
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		panic(err)
	}
}

func createServer() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc(serviceRoute+"/printmsg", printMsg).Methods("GET")
	r.HandleFunc(serviceRoute+"/alert", alert).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func printMsg(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(messages))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	}
}
func alert(w http.ResponseWriter, r *http.Request) {
	msg := message{}
	err := json.NewDecoder(r.Body).Decode(&msg)
	str := fmt.Sprintf(msg.ServiceName, msg.Reason, msg.TransactionType)
	messages = str
	_, err = w.Write([]byte(str))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	} else {
		log.Println(str)
	}
}
