package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	logError(err)

	for {
		conn, err := li.Accept()
		logError(err)
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	first := true
	var body string
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if first {
			split := strings.Fields(line)
			body = "You accessed " + split[1]
			first = false
		}
		if line == "" {
			break
		}
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func logError(e error) {
	if e != nil {
		log.Println(e)
	}
}
