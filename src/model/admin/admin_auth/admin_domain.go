package admin_auth_model

type adminAuthDomain struct {
	id                int
	email             string
	password          string
	encryptedPassword []byte
	salt              []byte
}

func (ad *adminAuthDomain) GetId() int{
	return ad.id
}

func (ad *adminAuthDomain) GetEmail() string {
	return ad.email
}

func (ad *adminAuthDomain) GetPassword() string {
	return ad.password
}

func (ud *adminAuthDomain) GetSalt() []byte {
	return ud.salt
}

func (ad *adminAuthDomain) GetEncryptedPassword() []byte {
	return ad.encryptedPassword
}

func (ad *adminAuthDomain) SetId(id int) {
	ad.id = id
}

func (ad *adminAuthDomain) SetEncryptedPassword(encryptedPassword []byte) {
	ad.encryptedPassword = encryptedPassword
}

func (ad *adminAuthDomain) SetSalt(salt []byte) {
	ad.salt = salt
}
