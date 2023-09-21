package email

import "net/mail"

func NewEmail(host,name,password string,port int) Email {
	return &email{
		host: host,
		name: name,
		password: password,
		port: port,
	}
}

type Email interface {
	SendEmail(to []mail.Address, subject, body string) error
}

type email struct {
	host string
	name string
	password string
	port int
}
