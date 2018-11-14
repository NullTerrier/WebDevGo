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
	defer li.Close()

	if err != nil {
		log.Println(err)
	}

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {

	defer conn.Close()

	request(conn)

	respond(conn)
}

func request(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	i := 0

	for scanner.Scan() {
		ln := scanner.Text()

		if i == 0 {
			elements := strings.Fields(ln)
			url := elements[1]
			fmt.Println("Your url: ", url)
		}

		if ln == "" {
			break
		}

		i++

	}

}

func respond(conn net.Conn) {
	body := `Whatever`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-LEngth: %d \r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
