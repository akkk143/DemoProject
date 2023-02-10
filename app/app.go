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

type message struct {
	ServiceName     string `json:"service_name"`
	Reason          string `json:"reason"`
	TransactionType string `json:"transaction_type"`
}

// Start starts the http server
func Start() {
	messages = append(messages, "Hello World!")
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
	r.HandleFunc(serviceRoute+"/customquery/alert", customeQueryAlert).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func printMsg(w http.ResponseWriter, r *http.Request) {
	for _, m := range messages {
		_, err := w.Write([]byte(m))
		w.Write([]byte("\n"))
		if err != nil {
			log.Printf("couldnt write response error [%s]\n", err)
		}
	}
}
func alert(w http.ResponseWriter, r *http.Request) {
	msg := message{}
	err := json.NewDecoder(r.Body).Decode(&msg)
	str := fmt.Sprintf(msg.ServiceName, msg.Reason, msg.TransactionType)
	if len(messages) > 2 {
		messages = make([]string, 0)
	}
	messages = append(messages, str)
	_, err = w.Write([]byte(str))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	} else {
		log.Println(str)
	}
}

type customeQuery struct {
	RuleQuery          string `json:"rule_query,omitempty"`
	Alerts             string `json:"alerts,omitempty"`
	ResultLink         string `json:"result_link,omitempty"`
	RuleDescription    string `json:"rule_description,omitempty"`
	ResponseActions    string `json:"response_actions,omitempty"`
	RuleReferences     string `json:"rule_references,omitempty"`
	ActionGroupName    string `json:"action_group_name,omitempty"`
	ActionSubgroupName string `json:"action_subgroup_name,omitempty"`
	RuleIndex          string `json:"rule_index,omitempty"`
}

type test struct {
	MyMessage string `json:"my_message"`
}

func customeQueryAlert(w http.ResponseWriter, r *http.Request) {
	msg := customeQuery{}
	err := json.NewDecoder(r.Body).Decode(&msg)
	log.Println(msg)
	//str := fmt.Sprintf("Alert : %s, RuleQuery : %s, RuleIndex : %s, ResultLink : %s, ResponseActions : %s", msg.Alerts, msg.RuleQuery, msg.RuleIndex,
	//	msg.ResultLink, msg.ResponseActions)
	str := fmt.Sprintf("ResponseActions : %s", msg.ResponseActions)
	if len(messages) > 2 {
		messages = make([]string, 0)
	}
	messages = append(messages, str)
	_, err = w.Write([]byte(str))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	} else {
		log.Println(str)
	}
}
