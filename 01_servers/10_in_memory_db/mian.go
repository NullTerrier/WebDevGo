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
	scanner := bufio.NewScanner(conn)

	defer conn.Close()

	data := make(map[string]string)

	for scanner.Scan() {
		cmd := scanner.Text()
		fs := strings.Fields(cmd)

		switch strings.ToUpper(fs[0]) {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "Your data at index %s : %s \n", k, v)
		case "SET":
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintf(conn, "INVALID COMMAND")
		}

	}
}
