package test

import (
	"fmt"
	"testing"
)

var f func()
var st *struct{}
var my_interface = []interface{}{f, st}

func TestNil(t *testing.T) {
	for i, v := range my_interface {
		if v != nil {
			fmt.Printf("CHECK != > %v, %T, %v\n", i, v, v)
		}
	}

	//不会打印, 因为interface含有类型以及值, 而这里v == nil只是值判断, 所以不生效
	for i, v := range my_interface {
		if v == nil {
			fmt.Printf("CHECK == > %v, %T, %v\n", i, v, v)
		}
	}
}
