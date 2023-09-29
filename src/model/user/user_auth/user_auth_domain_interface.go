package user_auth_model

type UserAuthDomainInterface interface {
	GetId() int64
	GetEmail() string
	GetName() string
	GetPassword() string
	GetEncryptedPassword() []byte
	GetSalt() []byte

	SetId(int64)
	SetPassword(string)
	SetEncryptedPassword([]byte)
	SetSalt([]byte)
	SetName(name string)
}

func NewUserSignUpDomain(email, name, password string) *userAuthDomain {
	return &userAuthDomain{
		email:    email,
		name:     name,
		password: password,
	}
}

func NewUserSignInDomain(email, password string) *userAuthDomain {
	return &userAuthDomain{
		email:    email,
		password: password,
	}
}

func NewUserDomainSendAuthEmail(id int64, email string, name string) *userAuthDomain {
	return &userAuthDomain{
		id:    id,
		email: email,
		name:  name,
	}
}

func NewUserDomainAuthEmail(id int64) *userAuthDomain {
	return &userAuthDomain{
		id: id,
	}
}
