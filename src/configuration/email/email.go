package email

import (
	"log"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"net/mail"
	"os"
	"strconv"
	"time"
)

func (e *email) SendEmail(id int64, to []mail.Address, subject, body string) error {
	err := e.configConn()
	if err != nil {
		return err
	}
	limitTime = time.Minute * 1
	user, err := e.handler(id)
	if err != nil {
		return err
	}
	quit := make(chan error, 1)
	go func() {
		e.conn(to, subject, body, quit)
	}()
	go func() {
		for {
			select {
			case q := <-quit:
				if q.Error() == "timeout" {
					delete(userEmail, id)
					close(quit)
					break
				}
				user.UserError = q
			default:
				return
			}
		}
	}()
	return nil
}

func (e *email) configConn() error {
	port,err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		logger.Error("Error loading email", err)
		log.Fatal("invalid port")
	}
	e.host = os.Getenv("EMAIL_HOST")
	e.name = os.Getenv("EMAIL_NAME")
	e.password = os.Getenv("EMAIL_PASSWORD")
	e.port = port
	return nil
}
