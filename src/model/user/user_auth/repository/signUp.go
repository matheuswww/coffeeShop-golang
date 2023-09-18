package user_auth_repository

import (
	"context"
	"errors"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	"time"

	"go.uber.org/zap"
)

func (ur userAuthRepository) SignUp(userAuthDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr {
	ctx,cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()
	logger.Info("Init createUser repository",zap.String("journey","CreateUser"))
	db := ur.databaseConnection
	query := "SELECT email from users WHERE email = ?"
	var existingEmail string
	rows,err := db.QueryContext(ctx,query,userAuthDomain.GetEmail())
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			logger.Error("Timeout expired",err,zap.String("journey","createUser"))
		} else {
			logger.Error("Error trying exec query",err,zap.String("journey","createUser"))
		}
		return rest_err.NewInternalServerError("database error")
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&existingEmail)
		if err != nil {
			logger.Error("Error trying scan query",err,zap.String("journey","createUser"))
			return rest_err.NewInternalServerError("database error")
		}
		logger.Error("Error trying insert user", errors.New("duplicated email"), zap.String("journey", "createUser"))
		return rest_err.NewConflictError("duplicated email")
	}
	query = "INSERT INTO users (email,name,password,salt,registration_date,last_access) VALUES (?, ?, ?, ?, ?, ?)"
	location := time.FixedZone("BRT", -3*60*60)
	timeStamp := time.Now().In(location).Format("2006-01-02 15:04:05")
	_,err = db.ExecContext(ctx,query,userAuthDomain.GetEmail(),userAuthDomain.GetName(),userAuthDomain.GetEncryptPassword(),userAuthDomain.GetSalt(),timeStamp,timeStamp)
	if err != nil {
		logger.Error("Error trying insert user",err,zap.String("journey","createUser"))
		return rest_err.NewInternalServerError("database error")
	}
	if err != nil {
		logger.Error("Error trying get id",err,zap.String("journey","createUser"))
		return rest_err.NewInternalServerError("database error")
	}
	logger.Info("USER INSERTED IN DATABASE")
	return nil
}