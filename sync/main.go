package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

//互斥锁
func sum() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
}

func write() {
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond)
	rwlock.Unlock()
	wg.Done()
}

func read() {
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	wg.Done()
}

func main() {
	/* wg.Add(2)
	go sum()
	go sum()
	wg.Wait()
	fmt.Println(x) */

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
