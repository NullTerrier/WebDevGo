package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	//must be closed
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panicln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Panicln(err)

		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)
	fmt.Fprintln(conn, "ello there!")

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "You have said: %s\n", ln)
	}

	defer conn.Close()

	fmt.Println("Code got here")
}
