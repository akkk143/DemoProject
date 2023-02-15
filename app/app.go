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

	m := body.(map[string]interface{})
	messages = append(messages, fmt.Sprint(body))
	//log.Println(m)
	api := ServiceData{}

	if data, ok := m["alert_data"]; ok {
		switch v := data.(type) {
		case map[string]interface{}:
			messages = append(messages, "data map[string]interface{}")
		case interface{}:
			messages = append(messages, "data interface{}")
			messages = append(messages, fmt.Sprint(v))
			s, _ := json.MarshalIndent(v, "", "\t")
			messages = append(messages, string(s))
			//for key, value := range v.(map[string]interface{}) {
			//	if key == "service" {
			//		b, err := json.Marshal(value)
			//		if err != nil {
			//			messages = append(messages, "service value marshaling failed")
			//		}
			//		service := model.Service{}
			//		err = json.Unmarshal(b, &service)
			//		if err != nil {
			//			messages = append(messages, "service value un marshaling failed")
			//		}
			//		api.ServiceName = service.Name
			//	} else if key == "url" {
			//		b, err := json.Marshal(value)
			//		if err != nil {
			//			messages = append(messages, "url value marshaling failed")
			//		}
			//		url := model.Url{}
			//		err = json.Unmarshal(b, &url)
			//		if err != nil {
			//			messages = append(messages, "url value un marshaling failed")
			//		}
			//		api.TransactionUrl = url.Full
			//	} else if key == "http" {
			//		b, err := json.Marshal(value)
			//		if err != nil {
			//			messages = append(messages, "http value marshaling failed")
			//		}
			//		http := model.Http{}
			//		err = json.Unmarshal(b, &http)
			//		if err != nil {
			//			messages = append(messages, "url value un marshaling failed")
			//		}
			//		api.StatusCode = http.Response.StatusCode
			//	} else if key == "transaction" {
			//		b, err := json.Marshal(value)
			//		if err != nil {
			//			messages = append(messages, "transaction value marshaling failed")
			//		}
			//		txn := model.Transaction{}
			//		err = json.Unmarshal(b, &txn)
			//		if err != nil {
			//			messages = append(messages, "transaction value un marshaling failed")
			//		}
			//		api.APIUrl = txn.Name
			//	}
			//	messages = append(messages, "out of the for loop")
			//}
		case map[string]string:
			messages = append(messages, "data map[string]string")
		case map[string]json.RawMessage:
			messages = append(messages, "data map[string]json.RawMessage")
		default:
			messages = append(messages, "data unknown")
		}

	} else {
		messages = append(messages, "data, ok := body.(map[string]interface{})[\"alert_data\"]; ok ")
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
