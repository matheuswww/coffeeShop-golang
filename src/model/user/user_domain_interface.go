package user_model

type UserDomainInterface interface {
	GetId() string
	GetEmail() string
	GetName() string
	GetPassword() string
	SetId(id int) error
	EncryptPassword() ([]byte,[]byte,error)
}

func NewUserDomain(email,name,password string) *userDomain{
	return &userDomain{
		email: email,
		name: name,
		password: password,
	}
}