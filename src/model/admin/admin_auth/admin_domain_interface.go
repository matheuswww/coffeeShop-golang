package admin_auth_model

type AdminAuthDomainInterface interface {
	GetId() int
	GetEmail() string
	GetPassword() string
	GetEncryptedPassword() []byte
	SetId(id int)
	SetEncryptedPassword(encryptedPassword []byte)
	GetSalt() []byte
	SetSalt(salt []byte)
}

func NewAdminAuthDomainInterface(email, password string) AdminAuthDomainInterface {
	return &adminAuthDomain{
		email:    email,
		password: password,
	}
}
