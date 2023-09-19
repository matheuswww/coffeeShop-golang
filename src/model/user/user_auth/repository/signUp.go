package user_auth_repository

import (
	"context"
	"fmt"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	"time"

	"strings"

	"go.uber.org/zap"
)

func (ur userAuthRepository) SignUp(userAuthDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Info("Init SignUp repository", zap.String("journey", "SignUp"))
	db := ur.databaseConnection
	query := `
		INSERT INTO users (email, name, password, salt, registration_date, last_access)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	location := time.FixedZone("BRT", -3*60*60)
	fmt.Println(userAuthDomain.GetSalt())
	timeStamp := time.Now().In(location).Format("2006-01-02 15:04:05")
	_, err := db.ExecContext(ctx, query, userAuthDomain.GetEmail(), userAuthDomain.GetName(), userAuthDomain.GetEncryptedPassword(), userAuthDomain.GetSalt(), timeStamp, timeStamp)
	if err != nil {
		logger.Error("Error trying insert user", err, zap.String("journey", "SignUp"))
		if strings.Contains(err.Error(), "Duplicate entry") {
			return rest_err.NewBadRequestError("email already exists")
		}
		return rest_err.NewInternalServerError("database error")
	}
	logger.Info("USER INSERTED IN DATABASE")
	return nil
}
