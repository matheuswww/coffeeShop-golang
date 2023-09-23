package user_auth_service

import (
	"errors"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

func (ud *userAuthDomainService) AuthEmail(userDomain user_auth_model.UserAuthDomainInterface, token string) *rest_err.RestErr {
	logger.Info("Init AuthEmail repository", zap.String("journey", "AuthEmail Repository"))
	tokenErr := verifyToken(token)
	if tokenErr != nil {
		logger.Error("Error trying AuthEmail", tokenErr, zap.String("journey", "AuthEmail Repository"))
		if tokenErr.Error() == "expired token" {
			return rest_err.NewUnauthorizeError("expired token")
		}
		return rest_err.NewBadRequestError("invalid token")
	}
	logger.Info("Init AuthEmail Service", zap.String("journey", "AuthEmail Service"))
	err := ud.userRepositroy.AuthEmail(userDomain)
	if err != nil {
		logger.Error("Error trying AuthEmail Service", err, zap.String("journey", "AuthEmail Service"))
		return err
	}
	return nil
}

func verifyToken(token string) error {
	secret := os.Getenv("JWTKEY")
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, errors.New("invalid assignature")
	})
	if err != nil {
		return err
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return errors.New("expired token")
			}
		} else {
			return errors.New("invalid token")
		}
		return errors.New("invalid jwt token")
	}
	return nil
}
