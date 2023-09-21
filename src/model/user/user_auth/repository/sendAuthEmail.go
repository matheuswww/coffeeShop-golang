package user_auth_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"matheuswww/coffeeShop-golang/src/configuration/email"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	"net/mail"
	"time"

	"go.uber.org/zap"
)

func (ur *userAuthRepository) SendAuthEmail(userDomain user_auth_model.UserAuthDomainInterface,token string) *rest_err.RestErr {
	logger.Info("Init SendAuthEmail Repository")
	db := ur.databaseConnection
	result,err := verifyAuthEmail(db,userDomain.GetId())
	if err != nil {
		logger.Error("Error trying verifyAuthEmail",err,zap.String("journey","SendAuthEmail"))
		if err.Error() == "user not found" {
			return rest_err.NewBadRequestError(err.Error())
		}
		return rest_err.NewInternalServerError("database error")
	}
	if result {
		logger.Error("Error SendAuthEmail verifyAuthEmail",err,zap.String("journey","SendAuthEmail"))
		return rest_err.NewBadRequestError("user already authenticated")
	} 
	logger.Info("Init user Repository",zap.String("journey","SendAuthEmail Repository"))
	to := []mail.Address {
		{ Name:userDomain.GetName(),Address: userDomain.GetEmail() },
	}
	email := email.NewEmail("172.17.0.1","","",1025)
	emailErr := email.SendEmail(to,"your link to authenticate your email",
		fmt.Sprintf(
			"<a href='%s/%s'>Autenticar meu email</a>",
			"http://localhost:8080/auth/email",
			token,
		),
	)
	if emailErr != nil {
		logger.Error("Error trying SendAuthEmail",emailErr,zap.String("journey","SendAuthEmail Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}

func verifyAuthEmail(database *sql.DB,id int64) (bool,error) {
	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 5))
	defer cancel()
	query := "SELECT authenticated FROM users WHERE id = ?"
	result,err := database.QueryContext(ctx,query,id)
	if err != nil {
		return false,err
	}
	var authenticated bool
	if result.Next() {
		if err := result.Scan(&authenticated); err != nil {
			return false,errors.New("error scanning result")
		}
	} else {		
		return true,errors.New("user not found")
	}
	return authenticated,err
}