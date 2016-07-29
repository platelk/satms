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
func (*ClientList) RegisterClient(val int) int {
	id := GenerateID()
	clients[id] = val
	return id
}

// GetClientList return an array with all the client registered
func (*ClientList) GetClientIDList() []int {
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

// Client is a convinient representation of a client registered in the service
type Client struct {
	// Id is an unique identifier for the client
	ID string

	conn *websocket.Conn
}

// internal counter
var uniqueID int

// GenerateID create an uniqueID for each new client
func GenerateID() int {
	uniqueID++
	return uniqueID
}

// Send a message to the Client
func (*Client) Send(msg []byte) {

}

// Unregister the client
func (*Client) Unregister() {

}
