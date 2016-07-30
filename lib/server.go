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

	http.HandleFunc("/message/send", func(w http.ResponseWriter, r *http.Request) {
		log.Println("On [/message/send] call")
		defer r.Body.Close()

		to, _ := strconv.ParseInt(r.URL.Query().Get("to"), 10, 32)
		from, _ := strconv.ParseInt(r.URL.Query().Get("from"), 10, 32)
		body := r.URL.Query().Get("body")

		log.Println("Send to client ", to)
		clientList.Get(int(to)).Send(&Message{int(from), int(to), body})
	})

	// WebSocket handler
	http.Handle("/client/connect", websocket.Handler(func(ws *websocket.Conn) {
		client := clientList.RegisterClient(ws)
		defer clientList.UnregisterClient(client)

		client.Listen()
	}))

	http.ListenAndServe(":8080", nil)
}
