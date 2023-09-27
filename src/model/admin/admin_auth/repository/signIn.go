package admin_auth_repository

import (
	"context"
	"crypto/subtle"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/mysql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_auth_model "matheuswww/coffeeShop-golang/src/model/admin/admin_auth"
	"matheuswww/coffeeShop-golang/src/model/util"
	"time"

	"go.uber.org/zap"
)

func (ar *adminAuthRepository) SignIn(adminDomain admin_auth_model.AdminAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init SignIn Repository", zap.String("journey", "SignIn Repository"))
	db, err := mysql.NewMysql().NewMysqlConnection()
	if err != nil {
		logger.Error("Error trying connect database", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 5))
	defer cancel()
	query := "SELECT id,password,salt FROM users WHERE email = ?"
	result, err := db.QueryContext(ctx, query, adminDomain.GetEmail())
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
	encrypt_err := util.EncryptUserPasswordWithSalt(adminDomain.GetPassword(), salt,
		func(hash, salt []byte) {
			adminDomain.SetEncryptedPassword(hash)
			adminDomain.SetSalt(salt)
		})
	if encrypt_err != nil {
		logger.Error("Error trying encrypt Password", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if subtle.ConstantTimeCompare(adminDomain.GetEncryptedPassword(), encryptedPassword) != 1 {
		logger.Error("Incorrect password or email", err, zap.String("journey", "SignIn Repository"))
		return rest_err.NewUnauthorizeError("Incorret password")
	}
	adminDomain.SetId(id)
	return nil
}
