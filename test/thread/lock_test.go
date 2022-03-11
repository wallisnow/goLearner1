package thread

import (
	"fmt"
	"sync"
	"testing"
)

var wg1 sync.WaitGroup
var lock sync.Mutex

func TestLoop(t *testing.T) {
	n := 0
	for i := 0; i < 99999; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			lock.Lock()
			n++
			lock.Unlock()
		}()
	}
	wg1.Wait()
	fmt.Println(n)
}
