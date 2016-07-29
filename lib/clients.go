// Package which will handle all the operations related to the client as :
//   - register a new client
//   - get a specific client
//   - ...
package satms

import "log"

// Client is a convinient representation of a client registered in the service
type Client struct {
	// Id is an unique identifier for the client
	ID string
}

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
