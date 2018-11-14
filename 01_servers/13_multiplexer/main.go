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

	returns := request(conn)

	fmt.Println(returns)

	switch returns[0] {
	case "GET":
		respondGet(conn, returns[1])
	case "Post":
		respondPost(conn, returns[1])
	default:
		log.Panicln("No such method")
	}

}

func request(conn net.Conn) []string {

	scanner := bufio.NewScanner(conn)
	var method string
	var url string
	var returns []string

	var i int

	for scanner.Scan() {
		ln := scanner.Text()

		if i == 0 {
			elements := strings.Fields(ln)
			method = elements[0]
			url = elements[1]
			fmt.Println("Your url: ", url)
		}

		if ln == "" {
			break
		}

		i++

	}

	returns = append(returns, method)
	returns = append(returns, url)

	return returns

}

func respondGet(conn net.Conn, url string) {

	// var message string

	// switch url {
	// case "/":
	// 	message = "Get, Hello world"
	// case "/marek":
	// 	message = "Get Hello Marek"
	// default:
	// 	message = "GET not defined"
	// }

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong> Hello World </strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-LEngth: %d \r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respondPost(conn net.Conn, url string) {

	var message string

	switch url {
	case "/":
		message = "POST, Hello world"
	case "/marek":
		message = "POST Hello Marek"
	default:
		message = "POST not defined"
	}

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>` + message + `</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-LEngth: %d \r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
