package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
