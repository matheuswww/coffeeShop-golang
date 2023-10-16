package email

import (
	"errors"
	"fmt"
	"time"
)

type userHandler struct {
	UserTime  time.Time
	UserError error
}

var (
	userEmail = make(map[string]*userHandler)
)

func (e *email) handler(id string) (*userHandler, error) {
	user, ok := userEmail[id]
	if ok {
		fmt.Println(user.UserError)
		if err := user.UserError; err != nil {
			return &userHandler{}, err
		}
		if time.Now().Add(limitTime).After(user.UserTime) {
			return &userHandler{}, errors.New("still sending email")
		} else {
			delete(userEmail, id)
		}
	}
	newUserHanlder := &userHandler{
		UserTime:  time.Now(),
		UserError: nil,
	}
	userEmail[id] = newUserHanlder
	return newUserHanlder, nil
}
