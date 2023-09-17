package user_model

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"

	"github.com/speps/go-hashids/v2"
	"go.uber.org/zap"
)

type userDomain struct {
	id string
	email string
	name string
	password string
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetId() string {
	return ud.id
}

func (ud *userDomain) SetId(id int) error {
	h,err := hashids.NewWithData(hashids.NewData())
	if err != nil {
		logger.Error("Error trying hash id",err,zap.String("journey","SetId"))
		return err
	}
	hash,err := h.Encode([]int{id})
	if err != nil {
		logger.Error("Error trying hash id",err,zap.String("journey","SetId"))
		return err
	}
	ud.id = hash
	return nil
}