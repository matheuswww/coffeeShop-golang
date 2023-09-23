package user_auth_repository

import (
	"context"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/mysql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	"time"

	"go.uber.org/zap"
)

func (ur *userAuthRepository) AuthEmail(userDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init AuthEmail repository", zap.String("journey", "AuthEmail Repository"))
	db, err := mysql.NewMysql().NewMysqlConnection()
	if err != nil {
		logger.Error("Error trying connect to database", err, zap.String("journey", "SignUp"))
		return rest_err.NewInternalServerError("database error")
	}
	defer db.Close()
	query := "UPDATE users SET authenticated = ? WHERE id = ?"
	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 5))
	defer cancel()
	_, err = db.ExecContext(ctx, query, true, userDomain.GetId())
	if err != nil {
		logger.Error("Error trying AuthEmail Repository", err, zap.String("journey", "AuthEmail"))
		return rest_err.NewInternalServerError("database error")
	}
	return nil
}
