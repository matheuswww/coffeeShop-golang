package email

import (
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

func (e *email) conn(to []mail.Address, subject, body string, quit chan error) error {
	smtpAddr := fmt.Sprintf("%s:%d", e.host, e.port)
	conn, err := net.Dial("tcp", smtpAddr)
	if err != nil {
		quit <- err
		return err
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(limitTime)); err != nil {
		err := errors.New("timeout")
		quit <- err
		return err
	}
	client, err := smtp.NewClient(conn, e.host)
	if err != nil {
		return err
	}
	from := "email@email.com"
	toAddresses := make([]string, len(to))
	for i, recipient := range to {
		toAddresses[i] = recipient.Address
	}
	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + strings.Join(toAddresses, ",") + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"Content-Type: text/html; charset=utf-8\r\n" +
			"\r\n" +
			body + "\r\n",
	)
	if err := client.Mail(from); err != nil {
		quit <- err
		return err
	}
	for _, recipient := range to {
		if err := client.Rcpt(recipient.Address); err != nil {
			quit <- err
			return err
		}
	}
	data, err := client.Data()
	if err != nil {
		quit <- err
		return err
	}
	_, err = data.Write(msg)
	if err != nil {
		quit <- err
		return err
	}
	data.Close()
	return nil
}
