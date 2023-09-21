package user_auth_model


type userAuthDomain struct {
	id int64
	email string
	name string
	password string
	encryptedPassword []byte
	salt []byte
}

func (ud *userAuthDomain) GetId() int64 {
	return ud.id
}

func (ud *userAuthDomain) GetEmail() string {
	return ud.email
}

func (ud *userAuthDomain) GetName() string {
	return ud.name
}

func (ud *userAuthDomain) GetPassword() string {
	return ud.password
}

func (ud *userAuthDomain) GetEncryptedPassword() []byte {
	return ud.encryptedPassword
}

func (ud *userAuthDomain) GetSalt() []byte {
	return ud.salt
}

func (ud *userAuthDomain) SetId(id int64) {
	ud.id = id
}

func (ud *userAuthDomain) SetPassword(password string) {
	ud.password = password 
}

func (ud *userAuthDomain) SetEncryptedPassword(encryptedPassword []byte) {
	ud.encryptedPassword = encryptedPassword
}

func (ud *userAuthDomain) SetSalt(salt []byte) {
	ud.salt = salt
}