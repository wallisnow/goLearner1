package Validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestNewUser(t *testing.T) {
	var ips = []string{"a", "b", "c", "e", "f", "g"}
	user := NewUser(ips)
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Printf("error %v\n", err)
	}
	fmt.Println(user)
}
