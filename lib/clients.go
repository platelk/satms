// Package satms will handle all the operations related to the client as :
//   - register a new client
//   - get a specific client
//   - ...
package satms

import (
	"log"

	"golang.org/x/net/websocket"
)

// client
var clients map[int]int

// ClientList will simplify management of mutiple client and basic action on this clients
type ClientList struct {
	clients map[int]*Client
}

// RegisterClient add a new client inside the list of already known client
func (clientList *ClientList) RegisterClient(ws *websocket.Conn) *Client {
	client := CreateClient(ws)
	clientList.clients[client.ID] = client
	return client
}

// GetClientIDList return an array with all the client registered
func (clientList *ClientList) GetClientIDList() []int {
	clientsID := make([]int, len(clientList.clients))
	i := 0
	for k := range clientList.clients {
		clientsID[i] = k
		i++
	}

	return clientsID
}

// CreateClientList create and initialize the necessary component for a ClientList
func CreateClientList() *ClientList {
	log.Println("Init clients data...")
	var cl ClientList

	cl.clients = make(map[int]*Client)

	return &cl
}

// Client is a convinient representation of a client registered in the service
type Client struct {
	// Id is an unique identifier for the client
	ID int

	conn *websocket.Conn
}

// internal counter
var uniqueID int

// GenerateID create an uniqueID for each new client
func GenerateID() int {
	uniqueID++
	return uniqueID
}

// CreateClient will allocate and initialize a new client
func CreateClient(ws *websocket.Conn) *Client {
	var newClient Client

	newClient.ID = GenerateID()
	newClient.conn = ws

	return &newClient
}

// Send a message to the Client
func (*Client) Send(msg []byte) {

}

// Unregister the client
func (*Client) Unregister() {

}
