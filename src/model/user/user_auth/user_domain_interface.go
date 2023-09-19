package user_auth_model

type UserAuthDomainInterface interface {
	GetId() int
	GetEmail() string
	GetName() string
	GetPassword() string
	GetEncryptedPassword() []byte
	GetSalt() []byte

	SetPassword(string)
	SetEncryptedPassword([]byte)
	SetSalt([]byte)
}

func NewUserAuthDomain(email,name,password string) *userAuthDomain{
	return &userAuthDomain{
		email: email,
		name: name,
		password: password,
	}
}