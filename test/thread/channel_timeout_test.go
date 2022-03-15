package thread

import (
	"fmt"
	"testing"
	"time"
)

var result = make(chan string)

func task(call func()) {
	time.Sleep(time.Second * 3)
	result <- "success"
	call()
}

func start() chan string {
	go task(callback)
	return result
}

func callback() {
	fmt.Println("callback")
}

func runTask() (interface{}, error) {
	start()
	select {
	case c := <-result:
		fmt.Println("task done")
		return c, nil
	case <-time.After(time.Second * 5):
		return nil, fmt.Errorf("\ntimeout")
	}
}

func TestTimeout(t *testing.T) {
	r, err := runTask()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
