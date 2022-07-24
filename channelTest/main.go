package main

import (
	"fmt"
	"time"
)

//生产者leak
func leak1() {
	ch := make(chan int)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- 100
	}()

	select {
	case res := <-ch:
		fmt.Printf("res: %d\n", res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout! exit...")
	}
}

//消费者leak
func leak2() {
	ch := make(chan int)

	go func() {
		for res := range ch {
			fmt.Printf("res: %d\n", res)
		}
		fmt.Println("done...")
	}()

	ch <- 1
	ch <- 2
	close(ch)
	time.Sleep(time.Second)
	fmt.Println("main goroutine done...")
}

func p1() {
	ch := make(chan int)
	done := make(chan struct{}, 1)

	go func() {
		<-time.After(2 * time.Second)
		fmt.Println("close2")
		close(ch)
		close(done)
	}()

	go func() {
		<-time.After(1 * time.Second)
		fmt.Println("close1")
		ch <- 1
		close(ch)
	}()

	<-done
}

func main() {
	p1()
}
