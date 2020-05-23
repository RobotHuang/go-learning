package main

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for index, arg := range os.Args {
		if index == 0 {
			continue
		} else {
			conn, err := connection(arg)
			defer conn.Close()
			if err != nil {
				log.Fatal(err)
				continue
			}
			wg.Add(1)
			go mustCopy(os.Stdout, conn)
		}
	}
	wg.Wait()
}

func connection(ip string) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", ip)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func mustCopy(dst io.Writer, src io.Reader) {
	defer wg.Done()
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}