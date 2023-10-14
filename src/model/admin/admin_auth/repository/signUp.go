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

func (ar *adminAuthRepository) SignUp(adminDomain admin_auth_model.AdminAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init SignUp repository", zap.String("journey", "SignUp Repository"))
	db := ar.database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `
		INSERT INTO admin (email, password, salt)
		VALUES (?, ?, ?)
	`
	result, err := db.ExecContext(ctx, query, adminDomain.GetEmail(),adminDomain.GetEncryptedPassword(), adminDomain.GetSalt())
	if err != nil {
		logger.Error("Error trying insert admin", err, zap.String("journey", "SignUp Repository"))
		if strings.Contains(err.Error(), "Duplicate entry") {
			return rest_err.NewBadRequestError("email already exists")
		}
		return rest_err.NewInternalServerError("database error")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error trying getting id", err, zap.String("journey", "SignUp Repository"))
		return rest_err.NewInternalServerError("server error,unable to generate session")
	}
	adminDomain.SetId(int(id))
	logger.Info("USER INSERTED IN DATABASE")
	return nil
}