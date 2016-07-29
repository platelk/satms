package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Println("Launching satms...")
	InitClients()
	LaunchServer()
}
