package user_model

type UserDomainInterface interface {
	GetId() int64
	GetEmail() string
	GetName() string
	GetPassword() string
	SetId(int64)
	EncryptPassword() ([]byte,[]byte,error)
}

func NewUserDomain(email,name,password string) *userDomain{
	return &userDomain{
		email: email,
		name: name,
		password: password,
	}
}