package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
			continue
		}

		// go handle(conn)
		io.WriteString(conn, "I see you connected\n")
		conn.Close()
	}
}

// func handle(conn net.Conn) {
// 	defer conn.Close()
// 	conn.Write([]byte("I see you connected\n"))
// }
