package test

import (
	"fmt"
	"testing"
	"time"
)

var result = make(chan string)

func task() {
	time.Sleep(time.Second * 3)
	result <- "success"
}

func start() chan string {
	go task()
	return result
}

func callback() {
	fmt.Println("callback")
}

func runTask() (interface{}, error) {
	taskRet := start()
	select {
	case c := <-taskRet:
		callback()
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
