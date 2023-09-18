package coockies

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SendCookie(c *gin.Context,id int) {
	session := sessions.Default(c)
	session.Set("id",id)
	session.Save()
}

func Store() cookie.Store {
	store := cookie.NewStore([]byte(os.Getenv("COOCKIEKEY")))
	store.Options(sessions.Options{
		Path: "/",
		MaxAge: 60 * 60 * 24,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteStrictMode,
	})
	return store
}