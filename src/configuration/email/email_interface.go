package email

import (
	"net/mail"
	"time"
)

var (
	limitTime time.Duration
)

func NewEmail() Email {
	return &email{}
}

type Email interface {
	SendEmail(id int64, to []mail.Address, subject, body string) error
	conn(to []mail.Address, subject, body string, quit chan error) error
	configConn() error
	handler(id int64) (*userHandler, error)
}

type email struct {
	host     string
	name     string
	password string
	port     int
}
