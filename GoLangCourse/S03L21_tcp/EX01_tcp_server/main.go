package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	logError(err)
	defer li.Close()

	for {
		conn, err := li.Accept()
		logError(err)
		io.WriteString(conn, "\nHello from TCP Server\n")
		fmt.Fprintln(conn, "How are you?")

		conn.Close()
	}

}

func logError(e error) {
	if e != nil {
		log.Println(e)
	}
}
