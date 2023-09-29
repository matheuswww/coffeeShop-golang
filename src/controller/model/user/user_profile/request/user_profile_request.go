package user_profile_request

type User_profile_request_addToCart struct {
	ProductId       string  `json:"productId" binding:"required"`
	ProductName     string  `json:"productName" binding:"required"`
	ProductPrice    float32 `json:"productPrice" binding:"required"`
	ProductQuantity int     `json:"productQuantity" binding:"required"`
}
