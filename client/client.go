package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

var (
	ip   string = "127.0.0.1"
	port int    = 2018
)

func AgentClient(ip string, port int, msgs []string) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln("Cannot connect to server:", err)
	}

	buff := make([]byte, 1024)
	for _, msg := range msgs {
		conn.Write([]byte(msg))
		log.Printf("Sent: %s", msg)
		n, _ := conn.Read(buff)
		log.Printf("Received: %s", buff[:n])
	}
}
