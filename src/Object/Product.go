package Object

type Product struct {
	Id   int    `myid:"1" youid:"2"`
	Name string `myname:"foo"`
}

func newProductWithAttrs(fs ...func(p *Product)) *Product {
	p := &Product{}
	for _, f2 := range fs {
		f2(p)
	}
	return p
}

func NewProductWithOneValue(f func(p *Product)) *Product {
	p := new(Product)
	f(p)
	return p
}

func NewProduct() *Product {
	return new(Product)
}

type ProductBuilder struct {
	id   int
	name string
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{}
}

func (builder *ProductBuilder) Id(id int) *ProductBuilder {
	builder.id = id
	return builder
}

func (builder *ProductBuilder) Name(name string) *ProductBuilder {
	builder.name = name
	return builder
}

func (builder *ProductBuilder) build() *Product {
	prod := NewProduct()
	prod.Id = builder.id
	prod.Name = builder.name
	return prod
}

func setId(id int) func(p *Product) {
	return func(p *Product) {
		p.Id = id
	}
}

func setName(name string) func(p *Product) {
	return func(p *Product) {
		p.Name = name
	}
}
