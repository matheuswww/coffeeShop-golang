package sessionCookie

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type sessionCookie struct {
	Id    string
	Email string
	Name  string
}

func Store() cookie.Store {
	store := cookie.NewStore([]byte(os.Getenv("COOCKIEKEY")))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24,
		HttpOnly: false,
		Secure:   true,
		Domain:   os.Getenv("DOMAIN"),
		SameSite: http.SameSiteNoneMode,
	})
	return store
}

func SendCoockie(c *gin.Context, id string, email string, name string) {
	sessions := sessions.Default(c)
	sessionCookie := sessionCookie{
		Id:    id,
		Email: email,
		Name:  name,
	}
	sessions.Set("id", sessionCookie.Id)
	sessions.Set("email", sessionCookie.Email)
	sessions.Set("name", sessionCookie.Name)
	sessions.Save()
}

func GetCookieValues(c *gin.Context) (sessionCookie, error) {
	sessions := sessions.Default(c)
	id := sessions.Get("id")
	email := sessions.Get("email")
	name := sessions.Get("name")
	if id != nil && email != nil && name != nil {
		return sessionCookie{
			Id:    id.(string),
			Email: email.(string),
			Name:  name.(string),
		}, nil
	}
	return sessionCookie{}, errors.New("invalid sessionCookie")
}
