package email

import (
	"errors"
	"net/mail"
	"os"
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
	mode := os.Getenv("MODE")
	if mode == "PROD" {
		return nil
	} else if mode == "DEV" {
		e.host = "172.17.0.1"
		e.name = ""
		e.password = ""
		e.port = 25
		return nil
	}
	return errors.New("invalid mode")
}
