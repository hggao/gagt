package main

import "log"

var (
	port int = 2018
)

func main() {

	AgentServer(port)

	log.Println("Server exit, bye!")
}
