package Object

type Product struct {
	Id   int    `myid:"1" youid:"2"`
	Name string `myname:"foo"`
}

func NewProduct() *Product {
	return &Product{}
}
