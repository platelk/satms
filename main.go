package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// client
var clients map[int]int

// internal counter
var uniqueID int

// GenerateID create an uniqueID for each new client
func GenerateID() int {
	uniqueID++
	return uniqueID
}

// RegisterClient add a new client inside the list of already known client
func RegisterClient(val int) int {
	id := GenerateID()
	clients[id] = val
	return id
}

// GetClientList return an array with all the client registered
func GetClientList() []int {
	clientsID := make([]int, len(clients))
	i := 0
	for k := range clients {
		clientsID[i] = k
		i++
	}

	return clientsID
}

// InitClients initialize the necessary component for client management and registration
func InitClients() {
	log.Println("Init clients data...")
	clients = make(map[int]int)
}

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
