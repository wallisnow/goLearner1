package String

import (
	"fmt"
	"strconv"
)

//定义自己的string, 其实还是返回自己的string
type String string

func From(str string) String {
	return String(str)
}

func FromInt(num int) String {
	return String(strconv.Itoa(num))
	//return String(fmt.Sprintf("%d", num))
}

func (this String) Len() int {
	return len(this)
}

func (this String) ForEach(f func(chars string)) {
	for i := 0; i < len(this); i++ {
		f(fmt.Sprintf("%c", this[i]))
	}
}
