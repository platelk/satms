package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// LaunchServer will launch a server that will handle client request
func LaunchServer() {
	log.Println("Launching http server...")

	http.HandleFunc("/client/register", func(w http.ResponseWriter, r *http.Request) {
		log.Println("On [/client/register] call")
		defer r.Body.Close()

		id := RegisterClient(42)
		w.Write([]byte(strconv.FormatInt(int64(id), 10)))
	})

	http.HandleFunc("/client/list", func(w http.ResponseWriter, r *http.Request) {
		log.Println("On [/client/list] call")
		defer r.Body.Close()

		w.Write([]byte(fmt.Sprint(GetClientList())))
	})

	http.HandleFunc("/message/send/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("On [/message/send] call")
		defer r.Body.Close()

	})

	http.ListenAndServe(":8080", nil)
}

func main() {
	log.Println("Launching satms...")
	InitClients()
	LaunchServer()
}
