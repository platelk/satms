package satms

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

func handleSocketMessage(clientList *ClientList, client *Client) {
	log.Print("Start listening socket message")
	for {
		select {
		case msg := <-client.MsgRecv:
			log.Print("Message receive !!")
			if msg.Topic == "myId" {
				client.Send(&Message{Topic: "myId", Body: strconv.FormatInt(int64(client.ID), 10)})
			} else if msg.Topic == "clientList" {
				client.Send(&Message{Topic: "clientList", Body: fmt.Sprint(clientList.GetClientIDList())})
			} else {
				clientList.Get(msg.To).Send(msg)
			}
		}
	}
}

// Handle websocket succesfull connection
func onClientConnect(ws *websocket.Conn) {
	defer func() { // Handle client disconnection
		ws.Close()
	}()

	log.Println("New client connection")
}

// InitServerRoute initialize all the routes the HTTP server will handle
func InitServerRoute() {
	log.Println("Launching http server...")
	clientList := CreateClientList(3)

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
		topic := r.URL.Query().Get("topic")

		log.Println("Send to client ", to)
		clientList.Get(int(to)).Send(&Message{topic, int(from), int(to), body})
	})

	// WebSocket handler
	http.Handle("/client/connect", websocket.Handler(func(ws *websocket.Conn) {
		client := clientList.RegisterClient(ws)
		defer clientList.UnregisterClient(client)

		go handleSocketMessage(clientList, client)
		client.Listen()
	}))
}

// LaunchServer will launch a server that will handle client request
func LaunchServer() {
	InitServerRoute()
	http.ListenAndServe(":4242", nil)
}
