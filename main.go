package main

import (
	"log"

	"github.com/platelk/satms/lib"
)

func main() {
	log.Println("Launching satms...")
	satms.InitClients()
	satms.LaunchServer()
}
