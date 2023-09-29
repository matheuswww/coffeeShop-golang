package user_profile_repository

import (
	"context"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/mysql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
	user_profile_model "matheuswww/coffeeShop-golang/src/model/user/user_profile"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (ur userProfileRepository) AddToCart(userProfileDomainSevice user_profile_model.UserProfileDomainInterface,productDomain product_model.ProductDomainInterface) *rest_err.RestErr {
	logger.Info("Init AddToCart Repository",zap.String("journey","AddToCart Repository"))
	db, err := mysql.NewMysql().NewMysqlConnection()
	if err != nil {
		logger.Error("Error trying connect database", err, zap.String("journey", "AddToCart Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	uuid,err := uuid.NewUUID()
	if err != nil {
		logger.Error("Error trying genereate uuid", err, zap.String("journey", "AddToCart Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	uuidString := uuid.String()
	query := "INSERT INTO cart (cart_id,user_id,product_id,product_name,quantity,price) VALUES(?, ?, ?, ?, ?, ?)"
	_,err = db.ExecContext(ctx,query,uuidString,userProfileDomainSevice.GetId(),productDomain.GetId(),productDomain.GetName(),productDomain.GetQuantity(),productDomain.GetPrice())
	if err != nil {
		logger.Error("Error trying insert cart", err, zap.String("journey", "AddToCart Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}