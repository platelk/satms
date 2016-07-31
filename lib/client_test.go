package satms

import (
	"io/ioutil"
	"log"
	"sort"
	"testing"
)

// TestCreateClientList will test if the object is correctly created from CreateClientList
func TestCreateClientList(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	var newClientList *ClientList

	newClientList = CreateClientList(1)
	if newClientList == nil {
		t.Error("Return nil object")
	}
	if newClientList.shards == nil {
		t.Error("Internal map not initialize")
	}
}

// TestRegister will test if the registering of an existing client works correctly
func TestRegister(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	newClientList := CreateClientList(1)
	newClient := CreateClient(nil)

	newClientList.Register(newClient)
	if newClientList.Get(newClient.ID) != newClient {
		t.Error("Wrong client registered")
	}
}

// TestRegisterClient will test if registering a non-existing client
func TestRegisterClient(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	newClientList := CreateClientList(1)
	newClient := newClientList.RegisterClient(nil)

	if len(newClientList.shards[0].clients) != 1 {
		t.Error("Client not correctly added in the internal map")
	}
	if newClientList.shards[0].clients[newClient.ID] != newClient {
		t.Error("Wrong client registered")
	}
}

// TestGet will test to retreive different previously registered client
func TestGet(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	newClientList := CreateClientList(1)
	var clientArray [5]*Client

	for i := 0; i < 5; i++ {
		clientArray[i] = newClientList.RegisterClient(nil)
	}

	for _, client := range clientArray {
		if newClientList.Get(client.ID) != client || newClientList.Get(client.ID).ID != client.ID {
			t.Error("Wrong client retreive in Get")
		}
	}
	if newClientList.Get(-42) != nil {
		t.Error("Nil is not returned when retreive an unknown client")
	}
}

// TestUnregisterClient will verify that on Unregister the client is correctly removed of the internal map
func TestUnregisterClient(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	newClientList := CreateClientList(3)
	var clientArray [5]*Client

	for i := 0; i < 5; i++ {
		clientArray[i] = newClientList.RegisterClient(nil)
	}

	for _, client := range clientArray {
		newClientList.UnregisterClient(client)
		if newClientList.Get(client.ID) != nil {
			t.Error("Client not correctly Unregister")
		}
	}
}

// TestGetClientIDList will test if the ClientList return all the ID of the actual registered client
func TestGetClientIDList(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	uniqueID = 0
	newClientList := CreateClientList(3)
	var clientArray [5]*Client
	var clientID []int

	for i := 0; i < 5; i++ {
		clientArray[i] = newClientList.RegisterClient(nil)
	}

	clientID = newClientList.GetClientIDList()
	sort.Ints(clientID)
	if len(clientID) != len(clientArray) {
		t.Error("List of client ID isn't the wright size")
	}
	for i, v := range []int{1, 2, 3, 4, 5} {
		if clientID[i] != v {
			t.Error("Wrong client ID in list")
		}
	}

	newClientList.UnregisterClient(clientArray[2])
	clientID = newClientList.GetClientIDList()
	sort.Ints(clientID)
	if len(clientID) != 4 {
		t.Error("List of client ID isn't the wright size after Unregister")
	}
	for i, v := range []int{1, 2, 4, 5} {
		if clientID[i] != v {
			t.Errorf("Missing client ID in list. %d != %d in %d position", clientID[i], v, i)
		}
	}
}

// TestCreateClient will verify if a new created client is correctly created
func TestCreateClient(t *testing.T) {
	newClient := CreateClient(nil)
	newClient2 := CreateClient(nil)

	if newClient == nil {
		t.Error("New client not correctly created")
	}
	if newClient.ID == newClient2.ID {
		t.Error("2 new create client have same ID !")
	}
}
