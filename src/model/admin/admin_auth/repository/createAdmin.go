package admin_auth_repository

import (
	"context"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_auth_model "matheuswww/coffeeShop-golang/src/model/admin/admin_auth"
	"strings"
	"time"

	"go.uber.org/zap"
)

func (ar *adminAuthRepository) CreateAdmin(adminDomain admin_auth_model.AdminAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init CreateAdmin repository", zap.String("journey", "CreateAdmin Repository"))
	db := ar.database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `
		INSERT INTO admin (email, password, salt)
		VALUES (?, ?, ?)
	`
	_, err := db.ExecContext(ctx, query, adminDomain.GetEmail(),adminDomain.GetEncryptedPassword(), adminDomain.GetSalt())
	if err != nil {
		logger.Error("Error trying insert admin", err, zap.String("journey", "CreateAdmin Repository"))
		if strings.Contains(err.Error(), "Duplicate entry") {
			return rest_err.NewBadRequestError("email already exists")
		}
		return rest_err.NewInternalServerError("database error")
	}
	logger.Info("USER INSERTED IN DATABASE")
	return nil
}