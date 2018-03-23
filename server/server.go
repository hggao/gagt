package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func AgentServer(port int) {

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	defer listener.Close()
	if err != nil {
		log.Fatalf("Listen on port %d failed, %s", port, err)
		os.Exit(1)
	}
	log.Printf("Listening on port: %d", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		log.Printf("Accepted connection: %v", conn)
		go connectionHandler(conn)
	}

}

func connectionHandler(conn net.Conn) {

	defer conn.Close()

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

	for {
		log.Printf("====== Start reading from connection.")
		n, err := r.Read(buf)
		if err == io.EOF {
			log.Printf("Got EOF, done with this connection.")
			return
		} else if err != nil {
			log.Fatalf("Receive data from connection failed: %s", err)
			return
		}
		if n == 0 {
			log.Printf("Got 0 bytes, read again.")
			continue
		}

		data := string(buf[:n])
		log.Println("Received:", data)
		w.Write([]byte(data))
		w.Flush()
		log.Printf("Sent: %s", data)
	}
}
