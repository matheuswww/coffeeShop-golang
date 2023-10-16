package user_auth_repository

import (
	"context"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	"time"

	"strings"

	"go.uber.org/zap"
)

func (ur userAuthRepository) SignUp(userAuthDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init SignUp repository", zap.String("journey", "SignUp Repository"))
	db := ur.database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `
		INSERT INTO users (id,email, name, password, salt, authenticated)
		VALUES (?,?, ?, ?, ?, ?)
	`
	_, err := db.ExecContext(ctx, query, userAuthDomain.GetEmail(), userAuthDomain.GetName(), userAuthDomain.GetEncryptedPassword(), userAuthDomain.GetSalt(), false)
	if err != nil {
		logger.Error("Error trying insert user", err, zap.String("journey", "SignUp Repository"))
		if strings.Contains(err.Error(), "Duplicate entry") {
			return rest_err.NewBadRequestError("email already exists")
		}
		return rest_err.NewInternalServerError("database error")
	}
	logger.Info("USER INSERTED IN DATABASE")
	return nil
}
