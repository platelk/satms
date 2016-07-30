// Package satms will handle all the operations related to the client as :
//   - register a new client
//   - get a specific client
//   - ...
package satms

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

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

// UnregisterClient will remove the client from the list of known client and will Unregister the client
func (clientList *ClientList) UnregisterClient(client *Client) {
	client.Unregister()
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

	return &ClientList{make(map[int]*Client)}
}

// Client is a convinient representation of a client registered in the service
type Client struct {
	// Id is an unique identifier for the client
	ID      int
	MsgRecv chan *Message

	conn      *websocket.Conn
	msgToSend chan *Message
	quitChan  chan bool
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
	newClient := &Client{GenerateID(), make(chan *Message), ws, make(chan *Message), make(chan bool)}

	log.Println("Create new client with ID = ", newClient.ID)

	return newClient
}

// Listen for incoming message and handle queue of message to send
func (client *Client) Listen() (err error) {
	for {
		select {
		case msg := <-client.msgToSend:
			client.write(msg)
		case <-client.quitChan:
			client.conn.Close()
			return nil
		default:
			var msg Message
			err := websocket.JSON.Receive(client.conn, &msg)
			if err == io.EOF {
				client.quitChan <- true
				return fmt.Errorf("Client %d is disconnected", client.ID)
			}
			client.MsgRecv <- &msg

		}
	}
}

// Internal write function
func (client *Client) write(msg *Message) {
	err := websocket.JSON.Send(client.conn, msg)
	if err != nil {
		client.Unregister()
	}
}

// Send a message to the Client
func (client *Client) Send(msg *Message) (err error) {
	select {
	case client.msgToSend <- msg:
	default:
		client.quitChan <- true
		return fmt.Errorf("Client %d is disconnected", client.ID)
	}
	return nil
}

// Unregister the client
func (client *Client) Unregister() {
	log.Println("Closing connection of Client: ", client.ID)
	client.quitChan <- true
}
