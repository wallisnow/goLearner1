package main

import (
	"base/src/String"
	"fmt"
)

func main() {
	s := String.From("mystring_type")
	fmt.Println(s)

	si := String.FromInt(123)
	fmt.Println(si)

	s.Len()
	fmt.Println(s.Len())

	//遍历字符串
	for i := 0; i < s.Len(); i++ {
		fmt.Println(i, " ", fmt.Sprintf("%c", s[i]))
	}

	s.ForEach(func(chars string) {
		fmt.Println(chars)
	})
}
