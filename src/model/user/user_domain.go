package user_model

type userDomain struct {
	id int64
	email string
	name string
	password string
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetId() int64 {
	return ud.id
}

func (ud *userDomain) SetId(id int64) {
	ud.id = id
}