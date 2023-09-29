package admin_product_model

type AdminProductDomainInterface interface {
	GetUUID() string
	GetName() string
	GetPrice() float32
	GetImg() []byte
	GetImgType() string
	GetStock() int64
	SetUUID(uuid string)
	SetImgType(imgType string)
}

func NewAdminProductModel(name string, price float32, img []byte, imgType string, stock int64) *adminProductDomain {
	return &adminProductDomain{
		name:    name,
		price:   price,
		img:     img,
		imgType: imgType,
		stock:   stock,
	}
}
