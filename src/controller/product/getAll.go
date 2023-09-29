package product_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (pc *productController) GetAll(c *gin.Context) {
	logger.Info("Init GetAll controller", zap.String("journey", "GetAll Controller"))
	var products []product_model.ProductDomainInterface
	err := pc.service.GetAll(&products)
	if err != nil {
		logger.Error("Error trying GetAll product Controller",err,zap.String("journey","GetAll Controller"))
		c.JSON(err.Code,err)
		return
	}
	var productJSONList []struct {
		ID    string  `json:"id"`
		Name  string  `json:"name"`
		Price float32 `json:"price"`
		Stock int     `json:"stock"`
	}
	for _, productDomain := range products {
		productJSON := struct {
				ID    string  `json:"id"`
				Name  string  `json:"name"`
				Price float32 `json:"price"`
				Stock int     `json:"stock"`
		}{
				ID:    productDomain.GetId(),
				Name:  productDomain.GetName(),
				Price: productDomain.GetPrice(),
				Stock: productDomain.GetStock(),
		}
		productJSONList = append(productJSONList, productJSON)
	}
	c.JSON(http.StatusOK,productJSONList)
}