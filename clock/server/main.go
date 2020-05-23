package main

import (
	"flag"
	"io"
	"net"
	"strings"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "port")
	flag.Parse()
	address := strings.Join([]string{"127.0.0.1", port}, ":")
	listen, err := net.Listen("tcp", address)
	if err != nil {
		println(err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		now := time.Now().Format("2006-01-02 15:04:05\n")
		addr := conn.LocalAddr()
		ip := addr.String()
		result := strings.Join([]string{ip, now}, "\t")
		_, err := io.WriteString(conn, result)
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
