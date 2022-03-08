package test

import (
	"fmt"
	"testing"
)

type Student struct{}

func NewStudent() *Student {
	return &Student{}
}

func (this *Student) change(i int) *Student {
	fmt.Println(i)
	return this
}

//链式调用中, defer只会在最后一个生效
//1
//2
//3
//10
//4
func TestDeferInChain(t *testing.T) {
	student := NewStudent()
	defer student.
		change(1).
		change(2).
		change(3).
		change(4)

	student.change(10)
}

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

//此时相当于
/*func TestDeferInLoopTest(t *testing.T) {
	i := 0
	defer func() {
		fmt.Println(i)
	}()
	i = 1
	defer func() {
		fmt.Println(i)
	}()
	i = 2
	defer func() {
		fmt.Println(i)
	}()
	i = 3
}*/

func TestDeferInloop(t *testing.T) {
	for i := 0; i < 3; i++ {
		//i := i
		defer func() {
			fmt.Println(i)
		}()
	}
}
