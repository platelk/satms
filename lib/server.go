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

	http.Handle("/client/connect", websocket.Handler(onClientConnect))

	http.ListenAndServe(":8080", nil)
}
