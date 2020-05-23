package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var tmp [128]byte
		len, err := reader.Read(tmp[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(tmp[:len]))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go process(conn)
	}

}
