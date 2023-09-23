package user_auth_repository

import (
	"context"
	"crypto/subtle"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/mysql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	user_auth_util "matheuswww/coffeeShop-golang/src/model/user/user_auth/util"
	"time"

	"go.uber.org/zap"
)

func (ur *userAuthRepository) SignIn(userDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init SignIn repository", zap.String("journey", "SignIn Repository"))
	db, err := mysql.NewMysql().NewMysqlConnection()
	if err != nil {
		logger.Error("Error trying connect to database", err, zap.String("journey", "SignUp"))
		return rest_err.NewInternalServerError("database error")
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 5))
	defer cancel()
	query := "SELECT id,password,salt FROM users WHERE email = ?"
	result, err := db.QueryContext(ctx, query, userDomain.GetEmail())
	if err != nil {
		logger.Error("Error trying SignIn user", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewInternalServerError("database error")
	}
	defer result.Close()
	var encryptedPassword, salt []byte
	var id int64
	if result.Next() {
		if err = result.Scan(&id, &encryptedPassword, &salt); err != nil {
			logger.Error("Error scanning result", err, zap.String("journey", "SignIn Repository"))
			return rest_err.NewInternalServerError("database error")
		}
	} else {
		logger.Error("Error email or password not found", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewUnauthorizeError("Email not registred")
	}
	encrypt_err := encryptUserPasswordWithSalt(userDomain, salt)
	if encrypt_err != nil {
		logger.Error("Error trying encrypt Password", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if subtle.ConstantTimeCompare(userDomain.GetEncryptedPassword(), encryptedPassword) != 1 {
		logger.Error("Incorrect password or email", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewUnauthorizeError("Incorret password")
	}
	userDomain.SetId(id)
	return nil
}

func encryptUserPasswordWithSalt(userDomain user_auth_model.UserAuthDomainInterface, salt []byte) *rest_err.RestErr {
	hash, _, err := user_auth_util.EncryptPassword(userDomain.GetPassword(), salt)
	if err != nil {
		logger.Error("Error tryin encrypt password", err, zap.String("journey", "SigIn Repository"))
		return rest_err.NewInternalServerError("database error")
	}
	userDomain.SetEncryptedPassword(hash)
	userDomain.SetSalt(salt)
	return nil
}
