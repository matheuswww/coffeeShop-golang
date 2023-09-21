package coockies

import (
	"errors"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type coockie struct {
	Id int64
	Email string
	Name string
}

func Store() cookie.Store {
	return cookie.NewStore([]byte(os.Getenv("COOCKIEKEY")))
}

func SendCoockie(c *gin.Context,id int64,email string,name string) {
	sessions := sessions.Default(c)
	coockie := coockie {
		Id: id,
		Name: name,
		Email: email,
	}
	sessions.Set("id",coockie.Id)
	sessions.Set("email",coockie.Email)
	sessions.Set("name",coockie.Name)
	sessions.Save()
}

func GetCookieValues(c *gin.Context) (coockie,error) {
	sessions := sessions.Default(c)
	id := sessions.Get("id")
	email := sessions.Get("email")
	name := sessions.Get("name")
	if id != nil && email != nil && name != nil {
		return coockie {
			Id: id.(int64),
			Email: email.(string),
			Name: name.(string),
		},nil
	}
	return coockie{},errors.New("invalid coockie")
}