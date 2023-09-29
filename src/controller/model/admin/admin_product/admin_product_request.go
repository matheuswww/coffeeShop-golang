package admin_product_request

import (
	"mime/multipart"
)

type AdminProductRequest struct {
	Name  string                `form:"name" binding:"required,min=10,max=150"`
	Price float32               `form:"price" binding:"required"`
	Img   *multipart.FileHeader `form:"img" binding:"required"`
	Stock int64                 `form:"stock" binding:"required"`
}
