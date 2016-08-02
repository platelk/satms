package satms

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"golang.org/x/net/websocket"
)

func BenchmarkNumberOfConnection(b *testing.B) {
	go LaunchServer(Config{})
	address := "localhost:4242"
	i := 0
	time.Sleep(2 * time.Second)
	for {
		ws, err := websocket.Dial(fmt.Sprintf("ws://%s/client/connect", address), "", fmt.Sprintf("http://%s/", address))
		client := CreateClient(ws)
		go func() {
			err := client.Listen()
			if err == nil {
				fmt.Printf("Connection closed")
				os.Exit(1)
			}
		}()

		if err != nil {
			fmt.Printf("Dial failed: %s\n", err.Error())
			os.Exit(1)
		}
		i++
		fmt.Printf("Nb : %d\n", i)
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	go LaunchServer(Config{})
	time.Sleep(2 * time.Second)
	os.Exit(m.Run())
}

func TestClientConnection(t *testing.T) {
	address := "localhost:4242"
	_, err := websocket.Dial(fmt.Sprintf("ws://%s/client/connect", address), "", fmt.Sprintf("http://%s/", address))

	if err != nil {
		t.Error("Client can't connect to server")
	}
}

func TestClientMyId(t *testing.T) {
	address := "localhost:4242"
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s/client/connect", address), "", fmt.Sprintf("http://%s/", address))

	if err != nil {
		t.Error("Client can't connect to server")
	}

	websocket.JSON.Send(ws, &Message{Topic: "myId"})
	var msg Message
	err = websocket.JSON.Receive(ws, &msg)
	if err != nil {
		t.Error("Error when receiving message")
	}
	if msg.Topic != "myId" {
		t.Error("Wrong topic receive")
	}
	if msg.Body != "9" {
		t.Error("Wrong return, expected: 9 got ", msg.Body)
	}
}

func TestClientClient(t *testing.T) {
	address := "localhost:4242"
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s/client/connect", address), "", fmt.Sprintf("http://%s/", address))

	if err != nil {
		t.Error("Client can't connect to server")
	}

	websocket.JSON.Send(ws, &Message{Topic: "clientList"})
	var msg Message
	err = websocket.JSON.Receive(ws, &msg)
	if err != nil {
		t.Error("Error when receiving message")
	}
	if msg.Topic != "clientList" {
		t.Error("Wrong topic receive")
	}
	if msg.Body != "[9 10 8]" {
		t.Error("Wrong return, expected: [9 10 8], got: ", msg.Body)
	}
}
