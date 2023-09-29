package product_model

type productDomain struct {
	id string
	name  string
	price float32
	stock int
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

func (pd *productDomain) GetStock() int {
	return pd.stock
}

func (pd *productDomain) SetId(id string) {
	pd.id = id
}

func (pd *productDomain) SetName(name string) {
	pd.name = name
}

func (pd *productDomain) SetPrice(price float32) {
	pd.price = price
}

func (pd *productDomain) SetStock(stock int) {
	pd.stock = stock
} 