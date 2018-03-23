package main

import "os"

func main() {
	if len(os.Args) > 1 {
		AgentClient(ip, port, os.Args[1:])
	} else {
		messages := []string{"Ping"}
		AgentClient(ip, port, messages)
	}
}
