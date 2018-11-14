package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "golang.org:80")

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}
