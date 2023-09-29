package product_model

type productDomain struct {
	id string
	name  string
	price float32
	quantity int
}

func (pd *productDomain) GetId() string {
	return pd.id
}

func (pd *productDomain) GetName() string {
	return pd.name
}

func (pd *productDomain) GetPrice() float32 {
	return pd.price
}

func (pd *productDomain) GetQuantity() int {
	return pd.quantity
}