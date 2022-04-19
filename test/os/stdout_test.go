package os

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestIO(t *testing.T) {
	_, err := io.WriteString(os.Stdout, "hello \n")
	if err != nil {
		return
	}

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!\r\n")
	w.Flush() // Don't forget to flush!
	//fmt.Fprint(os.Stdout, "hello, world!\r\n")

}
