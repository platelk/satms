package satms

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/websocket"
)

func BenchmarkNumberOfConnection(b *testing.B) {
	//go LaunchServer()
	address := "localhost:4242"
	i := 0
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
