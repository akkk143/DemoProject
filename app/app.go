package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const serviceRoute = "demo/project"

var srv *http.Server

// Start starts the http server
func Start() {
	createServer()
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		panic(err)
	}
}

func createServer() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc(serviceRoute+"/printmsg", printMsg).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func printMsg(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	}
}
