package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panicln(err)
	}

	defer li.Close()

	for {

		conn, err := li.Accept()

		if err != nil {
			log.Println(err)
		}

		go handle(conn)

	}

}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		log.Println("Connection timedout")

	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		li := scanner.Text()
		fmt.Println(li)

		fmt.Fprintf(conn, "I heard you say %s\n", li)
	}

	defer conn.Close()

	fmt.Println("CODE GOT HERE")
}
