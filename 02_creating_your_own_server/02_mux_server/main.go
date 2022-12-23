package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// TODO: Create a multiplexer so that your server responds to different URIs
	// and methods
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		m := fields[0]
		if m == "GET" {
			switch fields[1] {
			case "/":
				respond(conn, http.StatusOK, "home")
			case "/gallery":
				respond(conn, http.StatusOK, "gallery")
			default:
				respond(conn, http.StatusNotFound, "not found")
			}
			break
		}
		respond(conn, http.StatusBadRequest, "bad request")
		break
	}
}

func respond(conn net.Conn, status int, msg string) {
	l := "HTTP/1.1 " + strconv.Itoa(status) + " " + http.StatusText(status) + "\r\n"
	fmt.Fprintf(conn, l)
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(msg))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, msg)
}
