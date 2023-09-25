package admin_auth_model

type AdminAuthDomainInterface interface {
	GetId() int64
	GetEmail() string
	GetPassword() string
	GetEncryptedPassword() []byte
	SetId(id int64)
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
