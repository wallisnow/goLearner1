package thread

import (
	"fmt"
	"testing"
	. "time"
)

func Producer(queue chan int) {
	defer close(queue)
	for i := 0; i < 5; i++ {
		queue <- i
		Sleep(Second)
	}
}

func Consumer(queue chan int, waiter chan interface{}) {
	for num := range queue {
		fmt.Println(num)
	}
	waiter <- struct{}{}
}

func Consumer2(queue chan int) (r chan interface{}) {
	r = make(chan interface{})
	go func() {
		defer func() {
			r <- struct{}{}
		}()
		for num := range queue {
			fmt.Println(num)
		}
	}()
	return r
}

func TestConsumer(t *testing.T) {
	c := make(chan int)
	waiter := make(chan interface{})
	go Producer(c)
	go Consumer(c, waiter)
	<-waiter
}

func TestConsumer2(t *testing.T) {
	c := make(chan int)
	go Producer(c)
	ret := Consumer2(c)
	<-ret
}
