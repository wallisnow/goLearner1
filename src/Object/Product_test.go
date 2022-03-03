package Object

import (
	"fmt"
	"testing"
)

func TestNewProduct(t *testing.T) {
	productWithOnlyId := NewProductWithOneValue(setId(1))
	fmt.Println(productWithOnlyId)
	productWithAttrs := newProductWithAttrs(setId(1), setName("tester"))
	fmt.Println(productWithAttrs)

	productBuilder := NewProductBuilder()
	build := productBuilder.Id(123).Name("Milk").build()
	fmt.Println(build)
}
