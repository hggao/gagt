package comm

import "net"

func SendPacket(conn net.Conn, msg string) {
	conn.Write([]byte(msg))
}

func RecvPacket(conn net.Conn) (int, []byte) {
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	return n, buff
}
