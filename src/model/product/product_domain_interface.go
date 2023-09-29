package product_model

type ProductDomainInterface interface{
	GetId() string
	GetName() string
	GetPrice() float32
	GetQuantity() int
}

func NewProductDomain(id string,name string, price float32,quantity int) ProductDomainInterface {
	return &productDomain{ 
		id,
		name,
		price,
		quantity,
	}
}
