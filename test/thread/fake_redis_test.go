package thread

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var dict sync.Map

func Set(key string, value string, duration time.Duration) {
	dict.Store(key, value)
	time.AfterFunc(duration, func() {
		dict.Delete(key)
	})
}

func TestRedis(t *testing.T) {

	Set("1", "a", time.Second*2)
	Set("2", "b", time.Second*5)

	for {
		fmt.Println(dict.Load("1"))
		fmt.Println(dict.Load("2"))
		time.Sleep(time.Second)
	}

}
