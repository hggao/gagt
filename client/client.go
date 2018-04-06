package main

import (
	"github.com/hggao/gagt/lib"
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

	for _, msg := range msgs {
		comm.SendPacket(conn, msg)
		log.Printf("Sent: %s", msg)
		n, buff := comm.RecvPacket(conn)
		log.Printf("Received %d bytes: %s", n, buff)
	}
}
