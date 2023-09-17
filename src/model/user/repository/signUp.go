package user_repository

import (
	"context"
	"errors"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_model "matheuswww/coffeeShop-golang/src/model/user"
	"time"
	"go.uber.org/zap"
)

func (ur userRepository) SignUp(userDomain user_model.UserDomainInterface,hash,salt []byte) *rest_err.RestErr {
	ctx,cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()
	logger.Info("Init createUser repository",zap.String("journey","CreateUser"))
	db := ur.databaseConnection
	query := "SELECT email from users WHERE email = ?"
	var existingEmail string
	rows,err := db.QueryContext(ctx,query,userDomain.GetEmail())
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
	query = "INSERT INTO users (email,name,password,salt) VALUES (?, ?, ?, ?)"
	_,err = db.ExecContext(ctx,query,userDomain.GetEmail(),userDomain.GetName(),hash,salt)
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