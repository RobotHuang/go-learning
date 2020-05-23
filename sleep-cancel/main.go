package main

import (
	"sync"
	"time"
)

var done bool
var mu sync.Mutex

func main() {
	println("started")
	go period()
	time.Sleep(5 * time.Second)
	mu.Lock()
	done = true
	mu.Unlock()
	println("cancelled")
}

func period() {
	for {
		println("tick")
		time.Sleep(1 * time.Second)
		mu.Lock()
		if done {
			mu.Lock()
			return
		}
		mu.Unlock()
	}
}
