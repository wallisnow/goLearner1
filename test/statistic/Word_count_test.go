package statistic

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestCountLine(t *testing.T) {
	file, err := os.Open("../../go.sum")
	if err != nil {
		fmt.Println("error open")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
	}
	fmt.Println("count line: ", count)
}
