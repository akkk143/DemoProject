package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"log"
	"net/http"
	"strconv"
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
		w.Write([]byte("count : " + strconv.Itoa(count) + "\n"))
		if err != nil {
			log.Printf("couldnt write response error [%s]\n", err)
		}
	}
}

type ServiceData struct {
	ServiceName string
	APIUrl      string
	StatusCode  string
}

func customQueryAlert(w http.ResponseWriter, r *http.Request) {
	var body interface{}
	var bodyData = make(map[string]interface{}, 0)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		messages = append(messages, err.Error())
	}
	//messages = append(messages, fmt.Sprintf("%v", body))
	mapstructure.Decode(body, &bodyData)
	messages = append(messages, fmt.Sprintf("%v", bodyData))
	//var api = ServiceData{}
	//
	//data := bodyData["alert_data"].(map[string]interface{})
	//for i := 0; i < len(data); i++ {
	//	if data["fields"] != nil {
	//		valueMap := data["fields"].(map[string]interface{})
	//		for key, value := range valueMap {
	//			var str string
	//			switch value.(type) {
	//			case string:
	//				str = fmt.Sprint(value)
	//			case []interface{}:
	//				for _, v := range value.([]interface{}) {
	//					str = fmt.Sprint(v)
	//					break
	//				}
	//			}
	//			if key == "service.name" {
	//				api.ServiceName = str
	//			} else if key == "http.response.status_code" {
	//				api.StatusCode = str
	//			} else if key == "url.full" {
	//				api.APIUrl = str
	//			}
	//		}
	//	}
	//}
	//str := fmt.Sprint(api.APIUrl, " ", api.StatusCode, " ", api.ServiceName)
	//
	//if len(messages) > 2 {
	//	messages = make([]string, 0)
	//}
	//messages = append(messages, str)
	_, err = w.Write([]byte("Hello"))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	} else {
		log.Println("hello")
	}
}
