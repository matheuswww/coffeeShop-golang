package user_profile_repository

import (
	"context"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (ur userProfileRepository) AddToCart(userId string,productId string,quantity int) *rest_err.RestErr {
	logger.Info("Init AddToCart Repository", zap.String("journey", "AddToCart Repository"))
	db := ur.database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	uuid, err := uuid.NewUUID()
	if err != nil {
		logger.Error("Error trying genereate uuid", err, zap.String("journey", "AddToCart Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	uuidString := uuid.String()
	query := "INSERT INTO cart (cart_id,user_id,product_id,quantity) VALUES(?, ?, ?, ?)"
	_, err = db.ExecContext(ctx, query, uuidString,userId,productId,quantity)
	if err != nil {
		logger.Error("Error trying insert cart", err, zap.String("journey", "AddToCart Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}
