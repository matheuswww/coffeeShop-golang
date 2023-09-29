package product_model

type ProductDomainInterface interface {
	GetId() string
	GetName() string
	GetPrice() float32
	GetStock() int

	SetId(string)
	SetName(string)
	SetPrice(float32)
	SetStock(int)
}

func NewProductDomainGetAll() ProductDomainInterface {
	return &productDomain{}
}

func NewProductDomainService(id string, name string, price float32, stock int) ProductDomainInterface {
	return &productDomain{
		id,
		name,
		price,
		stock,
	}
}
