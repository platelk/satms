package satms

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

// Handle websocket succesfull connection
func onClientConnect(ws *websocket.Conn) {
	defer func() { // Handle client disconnection
		ws.Close()
	}()

	log.Println("New client connection")
}

// LaunchServer will launch a server that will handle client request
func LaunchServer() {
	log.Println("Launching http server...")
	clientList := CreateClientList()

	http.HandleFunc("/client/register", func(w http.ResponseWriter, r *http.Request) {
		log.Println("On [/client/register] call")
		defer r.Body.Close()

		client := clientList.RegisterClient(nil)

		w.Write([]byte(strconv.FormatInt(int64(client.ID), 10)))
	})

	http.HandleFunc("/client/list", func(w http.ResponseWriter, r *http.Request) {
		log.Println("On [/client/list] call")
		defer r.Body.Close()

		w.Write([]byte(fmt.Sprint(clientList.GetClientIDList())))
	})

	http.HandleFunc("/message/send/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("On [/message/send] call")
		defer r.Body.Close()

	})

	// WebSocket handler
	http.Handle("/client/connect", websocket.Handler(func(ws *websocket.Conn) {
		client := clientList.RegisterClient(ws)
		defer client.Unregister()
	}))

	http.ListenAndServe(":8080", nil)
}
