package main

import (
	"fmt"
	"time"
)

func say(message string) {
	fmt.Println(message)
}

func main() {
	time.Sleep(time.Second)
	selectChannel()
}

// goroutine
// go f(x, y, z)
// go func(){
// // function body here
// }()
// go aStruct.method()

// channel
// ch := make(chan T, size)
// buffered channel: size > 0
// range and close: v, ok := <-ch

func noBufferedChannel() {
	ch := make(chan string)
	go func() {
		ch <- "message"
		fmt.Println("done")
	}()
	time.Sleep(3 * time.Second)
	msg := <-ch
	fmt.Println(msg)
	time.Sleep(time.Second)
}

func bufferedChannel() {
	ch := make(chan string, 2)
	ch <- "msg1"
	ch <- "msg2"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func rangeAndCloseChannel() {
	ch := make(chan string)
	go func() {
		// ch <- "data"
		close(ch)
	}()

	// data, ok := <-ch
	// fmt.Println("data = ", data, "ok = ", ok)

	for data := range ch {
		fmt.Println("data = ", data)
	}
}

func selectChannel() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func() {
		ch1 <- "data"
	}()

	go func() {
		ch2 <- 1
	}()
	select {
	case data := <-ch1:
		fmt.Println("data from ch1 = ", data)
	case data2 := <-ch2:
		fmt.Println("data from ch2 = ", data2)
	default:
		fmt.Println("no data")
	}
}

// sync.Mutex, Lock() and Unlock()
mutex.Lock()
// do something
mutex.Unlock()

// sync.RWMutex

// Mutual Exclusion != Semaphore
