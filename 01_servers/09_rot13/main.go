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
	fmt.Fprintf(conn, "Please provide a string you want to rotate.")

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)

		fmt.Fprintf(conn, "%s - %s \n\n", ln, r)
	}

}

func rot13(bs []byte) []byte {

	var r13 = make([]byte, len(bs))

	for i, n := range bs {
		if n <= 109 {
			r13[i] = n + 13
		} else {
			r13[i] = n - 13
		}

	}

	return r13
}
