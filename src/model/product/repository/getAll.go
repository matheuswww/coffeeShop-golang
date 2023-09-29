package product_repository

import (
	"context"
	"errors"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/mysql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
	"time"

	"go.uber.org/zap"
)

func (pr *productRepository) GetAll(products *[]product_model.ProductDomainInterface) *rest_err.RestErr {
	logger.Info("Init GetAll repository", zap.String("journey", "GetAll Repository"))
	db, err := mysql.NewMysql().NewMysqlConnection()
	if err != nil {
		logger.Error("Error trying connect to database", err, zap.String("journey", "GetAll"))
		return rest_err.NewInternalServerError("database error")
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 5))
	defer cancel()
	query := "SELECT uuid,name,price,stock FROM products"
	result,err := db.QueryContext(ctx,query)
	if err != nil {
		logger.Error("Error trying GetAll products",err,zap.String("journey","GetAll repository"))
		return rest_err.NewInternalServerError("server error")
	}
	defer result.Close()
	for result.Next() {
		var id,name string
		var price float32
		var stock int 
		if result.Next() {
			if err := result.Scan(&id,&name,&price,&stock);err != nil {
				logger.Error("Error scanning result",err,zap.String("journey","GetAll Repository"))
				return rest_err.NewInternalServerError("server error")
			}
			product := product_model.NewProductDomainService(
				id,name,price,stock,
			)
			*products = append(*products,product)
		}
	}
	if len(*products) == 0 {
		logger.Error("Error no products found",errors.New("no products found"),zap.String("journey", "GetAll Repository"))
    return rest_err.NewNotFoundError("no products")
	}
	return nil
}