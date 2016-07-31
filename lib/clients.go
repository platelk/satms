// Package satms will handle all the operations related to the client as :
//   - register a new client
//   - get a specific client
//   - ...
package satms

import (
	"fmt"
	"io"
	"log"
	"sync"

	"golang.org/x/net/websocket"
)

// clientListShard is an internal type that allow the Client package to use Shard pattern to access client data
// Note: it's thread safe
type clientListShard struct {
	clients map[int]*Client
	mutex   *sync.Mutex
}

// Get return a client from its id
func (cls *clientListShard) Get(id int) *Client {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	return cls.clients[id]
}

// GetAllId return all the ID of the contained client inside the map
func (cls *clientListShard) GetAllID() []int {
	clientsID := make([]int, len(cls.clients))
	i := 0
	for k := range cls.clients {
		clientsID[i] = k
		i++
	}

	return clientsID
}

// Set add a client inside the map with the given id
func (cls *clientListShard) Set(id int, client *Client) *Client {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	cls.clients[id] = client

	return client
}

// Delete remove a client from the map
func (cls *clientListShard) Delete(id int) {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	delete(cls.clients, id)
}

// ClientList will simplify management of mutiple client and basic action on this clients
type ClientList struct {
	shards []*clientListShard
}

func (clientList *ClientList) shardHash(id int) int {
	return id % len(clientList.shards)
}

// RegisterClient add a new client inside the list of already known client
func (clientList *ClientList) RegisterClient(ws *websocket.Conn) *Client {
	client := CreateClient(ws)
	clientList.shards[clientList.shardHash(client.ID)].Set(client.ID, client)
	return client
}

// Register add an existing client inside the list of already known client
func (clientList *ClientList) Register(client *Client) *Client {
	clientList.shards[clientList.shardHash(client.ID)].Set(client.ID, client)
	return client
}

// Get a client by his ID
func (clientList *ClientList) Get(id int) *Client {
	return clientList.shards[clientList.shardHash(id)].Get(id)
}

// UnregisterClient will remove the client from the list of known client and will Unregister the client
func (clientList *ClientList) UnregisterClient(client *Client) {
	client.Unregister()
	clientList.shards[clientList.shardHash(client.ID)].Delete(client.ID)
}

// GetClientIDList return an array with all the client registered
func (clientList *ClientList) GetClientIDList() []int {
	var clientsID []int
	for _, k := range clientList.shards {
		for j := range k.clients {
			clientsID = append(clientsID, j)
		}
	}

	return clientsID
}

// CreateClientList create and initialize the necessary component for a ClientList
func CreateClientList(nbShards int) *ClientList {
	log.Println("Init clients data...")
	shards := make([]*clientListShard, nbShards)

	for i := 0; i < nbShards; i++ {
		shards[i] = &clientListShard{make(map[int]*Client), &sync.Mutex{}}
	}

	return &ClientList{shards}
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
	go func() {
		for {
			select {
			case <-client.quitChan:
				client.conn.Close()
			default:
				var msg Message
				err := websocket.JSON.Receive(client.conn, &msg)
				log.Println("Client ", client.ID, " receive : ", msg)
				if err == io.EOF {
					client.quitChan <- true
					return
				}
				msg.From = client.ID
				client.MsgRecv <- &msg
			}
		}
	}()
	for {
		select {
		case msg := <-client.msgToSend:
			msg.To = client.ID
			client.write(msg)
		case <-client.quitChan:
			client.conn.Close()
			return nil
		}
	}
}

// Internal write function
func (client *Client) write(msg *Message) {
	log.Println("client: write message ", msg)
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
	select {
	case client.quitChan <- true:
	default:
		log.Println("Unregister on already Unregistered client")
	}
}
