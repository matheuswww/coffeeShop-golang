package email

import (
	"context"
	"net/mail"
	"time"

	"gopkg.in/gomail.v2"
)

func (e *email) SendEmail(to []mail.Address, subject, body string) error {
	ctx,cancel := context.WithTimeout(context.Background(),time.Second * 3)
	defer cancel()
	emailSent := make(chan error,1)
	go func() {
		connect := gomail.NewDialer(e.host,e.port,e.name,e.password)
		message := gomail.NewMessage()
		message.SetHeader("From","email@gmail.com")
		for _,recipient := range to {
			message.SetHeader("To",recipient.Address)
		}
		message.SetHeader("Subject",subject)
		message.SetBody("text/html",body)
		err := connect.DialAndSend(message)
		emailSent <- err
	}()
	select {
	case err := <-emailSent:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}