package user_auth_service

import (
	"fmt"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

func (ud userAuthDomainService) SendAuthEmail(userDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr {	
	logger.Info("Init SendAuthEmail service",zap.String("journey","SendAuthEmail Service"))
	token,err := generateToken(userDomain)
	if err != nil {
		logger.Error("Error trying generate token",err,zap.String("journey","SendAuthEmail service"))
		return err
	}
	err = ud.userRepositroy.SendAuthEmail(userDomain,token)
	if err != nil {
		logger.Error("Error trying auth user",err,zap.String("journey","SendAuthEmail Service"))
		return err
	}
	return nil
}

func generateToken(userDomain user_auth_model.UserAuthDomainInterface) (string,*rest_err.RestErr) {
	secret := os.Getenv("JWTKEY")
	fmt.Println(time.Now().Unix())
	claims := jwt.MapClaims{
		"exp": int64(time.Now().Add(time.Minute * 2).Unix()),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"sub": userDomain.GetId(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,claims)
	tokenString,err := token.SignedString([]byte(secret))
	if err != nil {
		return "",rest_err.NewInternalServerError("server error")
	}
	return tokenString,nil
}