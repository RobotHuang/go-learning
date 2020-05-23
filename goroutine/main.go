package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

var jobChan chan int64 = make(chan int64, 100)
var resultChan chan int64 = make(chan int64, 100)

func generate(job chan<- int64) {
	for {
		x := rand.Int63()
		job <- x
		time.Sleep(time.Millisecond * 500)
	}
}

func calc(result chan<- int64, job <-chan int64) {
	for {
		value := <-job
		var sum int64 = 0
		for value > 0 {
			sum += value % 10
			value /= 10
		}
		result <- sum
	}
}

func main() {
	/* go hello() */
	/* for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	} */
	/* for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	} */
	/* for i := 0; i < 100; i++ {
		wg.Add(1)
		go hello(i)
	} */
	//fmt.Println("main")
	//time.Sleep(time.Second)
	/* runtime.GOMAXPROCS(2)
	wg.Add(2)
	go a()
	go b() */
	wg.Add(1)
	go generate(jobChan)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go calc(resultChan, jobChan)
	}
	for i := range resultChan {
		fmt.Println(i)
	}
	wg.Wait()
}
