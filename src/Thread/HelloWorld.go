package Thread

import (
	"fmt"
)

func HelloChan() {
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		go func(n int) {
			ch <- n * 2
		}(i)
	}

	for i := 0; i < cap(ch); i++ {
		fmt.Println("====> ", <-ch)
	}

}

// 场景1
func WriteNoBufChWithGoroutine() {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 1
	}(ch)
	<-ch
}

// 场景2 会报错
func WriteNoBufCh() {
	ch := make(chan int)
	func(ch chan int) {
		fmt.Printf("will hava exception: fatal error: all goroutines are asleep - deadlock!\n")
		ch <- 1
	}(ch)
	<-ch
}
