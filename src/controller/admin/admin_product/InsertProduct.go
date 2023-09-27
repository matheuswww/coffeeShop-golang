package admin_product_controller

import (
	"errors"
	"io"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	"matheuswww/coffeeShop-golang/src/configuration/validation"
	admin_product_request "matheuswww/coffeeShop-golang/src/controller/model/admin/admin_product"
	"matheuswww/coffeeShop-golang/src/controller/routes/coockies"
	admin_product_model "matheuswww/coffeeShop-golang/src/model/admin/admin_product"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (ac *adminProductController) InsertProduct(c *gin.Context) {
	logger.Info("Init InsertProduct Controller",zap.String("journey","InsertProduct Controller"))
	value,coockieErr := coockies.GetCookieValues(c)
	if coockieErr != nil {
		logger.Error("Error invalid coockie", coockieErr, zap.String("journey", "InsertProduct Controller"))
		restErr := rest_err.NewBadRequestError(coockieErr.Error())
		c.JSON(restErr.Code, restErr)
		return
	}
	if value.Name != "admin" {
		c.Status(401)
		return
	}
	var adminProductRequest admin_product_request.AdminProductRequest
	if err := c.ShouldBind(&adminProductRequest); err != nil {
		logger.Error("Error trying InsertProduct Controller",err,zap.String("journey","InsertProduct Controller"))
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code,restErr)
		return
	}
	image,imageType,imageErr := ac.readImage(*adminProductRequest.Img)
	if imageErr != nil {
		if imageErr.Error() == "invalid image type" {
			logger.Error("Error reading image",imageErr,zap.String("journey","InsertProduct Controller"))
			restErr := validation.ValidateError(imageErr)
			c.JSON(restErr.Code,restErr)
			return
		} else {
			logger.Error("Error reading image",imageErr,zap.String("journey","InsertProduct Controller"))
			restErr := validation.ValidateError(errors.New("the image could not be read, the file may be corrupted"))
			c.JSON(restErr.Code,restErr)
			return
		}
	} 
	domain := admin_product_model.NewAdminProductModel(
		adminProductRequest.Name,
		adminProductRequest.Price,
		image,
		imageType,
		adminProductRequest.Stock,
	)
	err := ac.service.InsertProduct(domain)
	if err != nil {
		logger.Error("Error trying InsertProduct Controller",err,zap.String("journey","InsertProduct Controller"))
		c.JSON(err.Code,err)
		return
	}
	c.Status(http.StatusCreated)
}

func (ac adminProductController) readImage(fileHeader multipart.FileHeader) ([]byte,string,error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil,"",err
	}
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil,"",err
	}
	var fileType string
	switch http.DetectContentType(fileBytes) {
		case "image/jpeg":
			fileType = "image/jpeg"
		case "image/jpg":
			fileType = "image/png"
		case "image/webp":
			fileType = "image/webp"
		default:
			fileType = ""
	}
	if fileType == "" {
		return nil,"",errors.New("invalid image type")
	} else {
		fileType = strings.TrimPrefix(fileType,"image/")
	}
	return fileBytes,fileType,nil
}