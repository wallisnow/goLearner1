package thread

import (
	"fmt"
	"golang.org/x/sys/unix"
	"strconv"
	"sync"
	"testing"
	. "time"
)

var wg sync.WaitGroup
var batchCh = make(chan struct{}, 10)

func run(str string) {
	defer wg.Done()
	//写到10个时阻塞
	batchCh <- struct{}{}
	fmt.Printf("%v %s \n", unix.Gettid(), str)
	Sleep(Second)
	<-batchCh
}

func TestParallel(t *testing.T) {
	for i := 0; i < 10000; i++ {
		//t.Skip("Skipping testing")
		wg.Add(1)
		go run("Runner No." + strconv.Itoa(i) + " finished")
	}
	wg.Wait()
}

func TestWaitGroup(t *testing.T) {
	var wgp sync.WaitGroup
	for i := 0; i < 10; i++ {
		wgp.Add(1)
		go func(i int) {
			defer wgp.Done()
			fmt.Println("Runner No.", strconv.Itoa(i), " finished")
			Sleep(Second)
		}(i)
	}
	wgp.Wait()
}
