package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")

	fmt.Println("Started")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	fmt.Println("Started")

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Println(err)
			continue

		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	fmt.Println("Started")

	for scanner.Scan() {

		fmt.Println("Running")
		ln := scanner.Text()

		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say : %s\n", ln)
	}

	defer conn.Close()

	fmt.Println("Code got here.")
}
