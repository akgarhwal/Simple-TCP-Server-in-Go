package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
	"time"
)

func getNameFromRequest(request string) string {

	/*
		Request from client:  GET /?name=Abhinesh HTTP/1.1
		Host: localhost:8080
		User-Agent: curl/7.86.0
		Accept: *\/*
	*/

	// Split the request by newline characters
	lines := strings.Split(request, "\n")
	// Get the first line, which contains the URL
	firstLine := lines[0]
	// Split the first line by spaces
	parts := strings.Split(firstLine, " ")
	// Get the second part, which is the URL path and query
	pathAndQuery := parts[1]
	// Parse the URL
	u, err := url.Parse(pathAndQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Get the query parameters
	q := u.Query()
	// Get the value of name parameter
	name := q.Get("name")

	return name
}

func getHttpResponse(name string) string {
	return "HTTP/1.1 200 OK \r\n\r\nHello " + name + "!\r\n"
}

func handleConnection(conn net.Conn) {
	log.Println("Processing the request")

	buf := make([]byte, 1024)
	// Read Connection in Buffer
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	request := string(buf)
	fmt.Println("Request from client: ", request)
	name := getNameFromRequest(request)

	time.Sleep(10 * time.Second)

	response := getHttpResponse(name)
	// Write back to Connection
	conn.Write([]byte(response))

	// Close the Connection
	conn.Close()
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("TCP Server started on port:8080\n ")

	for {
		log.Println("Waiting for a cleint to connect")

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Cleint connected")

		go handleConnection(conn)
	}
}
