package admin_product_repository

import (
	"context"
	"encoding/base64"
	"fmt"
	"go.uber.org/zap"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_product_model "matheuswww/coffeeShop-golang/src/model/admin/admin_product"
	"os"
	"time"
)

func (ar *adminProductRepository) InsertProduct(AdminProductDomain admin_product_model.AdminProductDomainInterface) *rest_err.RestErr {
	logger.Info("Init InsertProduct Repository", zap.String("journey", "InsertProduct Repository"))
	db := ar.database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	query := "INSERT INTO products (uuid, name, price, stock) VALUES (?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, query, AdminProductDomain.GetUUID(), AdminProductDomain.GetName(), AdminProductDomain.GetPrice(), AdminProductDomain.GetStock())
	if err != nil {
		logger.Error("Error trying insert user", err, zap.String("journey", "InsertProduct Repository"))
		return rest_err.NewInternalServerError("database error")
	}
	err = ar.upload(AdminProductDomain)
	if err != nil {
		logger.Error("Error trying upload image", err, zap.String("journey", "InsertProduct Repository"))
		return rest_err.NewInternalServerError("the product was inserted,but the image not uploaded")
	}
	return nil
}

func (ar *adminProductRepository) upload(AdminProductDomain admin_product_model.AdminProductDomainInterface) error {
	root, errRoot := os.Getwd()
	if errRoot != nil {
		return errRoot
	}
	uuidBytes := []byte(AdminProductDomain.GetUUID())
	dst, err := os.Create(fmt.Sprintf("%s/src/public/img/%s.%s", root, base64.StdEncoding.EncodeToString(uuidBytes), AdminProductDomain.GetImgType()))
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = dst.Write(AdminProductDomain.GetImg())
	if err != nil {
		return err
	}
	return nil
}
