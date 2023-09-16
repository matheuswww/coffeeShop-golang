package user_view

import (
	user_response "matheuswww/coffeeShop-golang/src/controller/model/user/response"
	user_model "matheuswww/coffeeShop-golang/src/model/user"
	"strconv"
)

func ConvertDomainToResponse(userDomain user_model.UserDomainInterface) user_response.User_response {
	Idstr := strconv.FormatInt(userDomain.GetId(),10)
	return user_response.User_response{
		Id: Idstr,
		Name: userDomain.GetEmail(),
		Email: userDomain.GetEmail(),
	}
}