package admin_product_model

type adminProductDomain struct {
	uuid    string
	name    string
	price   float32
	img     []byte
	imgType string
	stock   int64
}

func (ad *adminProductDomain) GetUUID() string {
	return ad.uuid
}

func (ad *adminProductDomain) GetName() string {
	return ad.name
}

func (ad *adminProductDomain) GetPrice() float32 {
	return ad.price
}

func (ad *adminProductDomain) GetStock() int64 {
	return ad.stock
}

func (ad *adminProductDomain) GetImg() []byte {
	return ad.img
}

func (ad *adminProductDomain) GetImgType() string {
	return ad.imgType
}

func (ad *adminProductDomain) SetUUID(uuid string) {
	ad.uuid = uuid
}

func (ad *adminProductDomain) SetImgType(imgType string) {
	ad.imgType = imgType
}
