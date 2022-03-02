package Object

import "fmt"

type ProdService struct{}

func NewProdService() *ProdService {
	return &ProdService{}
}

func (this *ProdService) Save(p interface{}) IService {
	fmt.Println("save product")
	return this
}

func (this *ProdService) List() IService {
	fmt.Println("list products")
	return this
}
