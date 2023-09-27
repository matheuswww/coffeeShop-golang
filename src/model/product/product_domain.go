package product_model

type productDomain struct {
	name  string
	price float32
	image []byte
}

func (pd *productDomain) GetName() string {
	return pd.name
}

func (pd *productDomain) GetPrice() float32 {
	return pd.price
}

func (pd *productDomain) GetImage() []byte {
	return pd.image
}
