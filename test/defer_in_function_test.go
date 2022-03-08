package test

import (
	"fmt"
	"testing"
)

//这里因为, 参数在传入函数时其实已经是确定的值, 所以defer 是不会有任何改变
func TestDeferStillReturn10(t *testing.T) {
	i := 10
	defer fmt.Println(i)
	i++
}

//如果传入指针, 这样就避免了上面的问题, 因为此时我们传入指针, i++时, 只是值改变, 但我们传入的是指针, 那么依然可以拿到其改变后的值
func TestDeferStillReturn11(t *testing.T) {
	i := 10
	var p *int
	p = &i
	defer func(p *int) {
		fmt.Println(*p)
	}(p)
	i++
}
