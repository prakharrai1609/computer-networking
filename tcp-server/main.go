package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func connect(conn net.Conn) {
	buf := make([]byte, 1024)

	n, err := conn.Read(buf)

	if err != nil {
		log.Fatal(err)
	}

	requestData := buf[:n]
	fmt.Println("Received data from client:", string(requestData))

	time.Sleep(5 * time.Second)

	response := "HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"
	conn.Write([]byte(response))

	conn.Close()
}

func main() {
	fmt.Println("Server started...")
	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go connect(conn)
	}
}
