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

	for i, v := range my_interface {
		if v == nil {
			fmt.Printf("CHECK == > %v, %T, %v\n", i, v, v)
		}
	}
}
