package user_auth_repository

import (
	"context"
	"crypto/subtle"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	"matheuswww/coffeeShop-golang/src/model/util"
	"time"

	"go.uber.org/zap"
)

func (ur *userAuthRepository) SignIn(userDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init SignIn repository", zap.String("journey", "SignIn Repository"))
	db := ur.database
	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 5))
	defer cancel()
	query := "SELECT id,password,salt,name FROM users WHERE email = ?"
	result, err := db.QueryContext(ctx, query, userDomain.GetEmail())
	if err != nil {
		logger.Error("Error trying SignIn user", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewInternalServerError("database error")
	}
	defer result.Close()
	var encryptedPassword, salt []byte
	var id int64
	var name string
	if result.Next() {
		if err = result.Scan(&id, &encryptedPassword, &salt, &name); err != nil {
			logger.Error("Error scanning result", err, zap.String("journey", "SignIn Repository"))
			return rest_err.NewInternalServerError("database error")
		}
	} else {
		logger.Error("Error email or password not found", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewUnauthorizedError("Email not registred")
	}
	encrypt_err := util.EncryptUserPasswordWithSalt(userDomain.GetPassword(), salt,
		func(hash, salt []byte) {
			userDomain.SetEncryptedPassword(hash)
			userDomain.SetSalt(salt)
		})
	if encrypt_err != nil {
		logger.Error("Error trying encrypt Password", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if subtle.ConstantTimeCompare(userDomain.GetEncryptedPassword(), encryptedPassword) != 1 {
		logger.Error("Incorrect password or email", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewUnauthorizedError("Incorret password")
	}
	userDomain.SetId(id)
	userDomain.SetName(name)
	return nil
}
